function base64UrlToBuffer(data: string): ArrayBuffer {
  return Uint8Array.fromBase64(data, { alphabet: "base64url" }).buffer;
}

function bufferToBase64Url(buffer: ArrayBuffer): string {
  return new Uint8Array(buffer).toBase64({ alphabet: "base64url" });
}

type ExistingCredential = { id: string };

async function getCredentialCreationOptions(): Promise<CredentialCreationOptions> {
  const { publicKey } = await fetch("/webauthn/register/start", {
    credentials: "include",
  }).then((resp) => resp.json());

  return {
    publicKey: {
      ...publicKey,
      challenge: base64UrlToBuffer(publicKey.challenge),
      user: {
        ...publicKey.user,
        id: base64UrlToBuffer(publicKey.user.id),
      },
      excludeCredentials: publicKey.excludeCredentials?.map((c: ExistingCredential) => ({
        ...c,
        id: base64UrlToBuffer(c.id),
      })),
    },
  };
}

export async function finalizeCredentialRegistration(credential: PublicKeyCredential) {
  const response = await fetch("/webauthn/registration/finalize", {
    body: JSON.stringify(serializeCredential(credential)),
    headers: {
      "Content-Type": "application/json",
    },
    method: "POST",
  });
  return response.json();
}

function serializeCredential(credential: PublicKeyCredential) {
  const response = credential.response as AuthenticatorAssertionResponse;
  return {
    id: credential.id,
    rawId: bufferToBase64Url(credential.rawId),
    type: credential.type,
    response: {
      authenticatorData: bufferToBase64Url(response.authenticatorData),
      clientDataJSON: bufferToBase64Url(response.clientDataJSON),
      signature: bufferToBase64Url(response.signature),
      userHandle: response.userHandle ? bufferToBase64Url(response.userHandle) : undefined,
    },
    clientExtensionResults: credential.getClientExtensionResults(),
  };
}

function initWebauthnRegistration() {
  const form = document.querySelector("form#webauthn-registration");
  if (!form) return;

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const options = await getCredentialCreationOptions();
    const credential = await navigator.credentials.create(options);
    if (!credential) return;
  });
}

initWebauthnRegistration();
