import React, { Component } from "react";
import "./Login.css";

class Login extends Component {
  render() {
    return (
      <html class="html__responsive html__unpinned-leftnav">
      <div class= "container">
        <div id="content" style={{backgroundColor: 'white'}} class="grid grid__center">
          <div class="bgr-gr">

            <div
              id="openid-buttons"
              class="mx-auto grid grid__fl1 fd-column gs8 gsy mb16 wmx3"
            >
              <button
                class="grid--cell s-btn s-btn__icon s-btn__google bar-md ba bc-black-100"
                data-provider="google"
                data-oauthserver="https://accounts.google.com/o/oauth2/auth"
                data-oauthversion="2.0"
              >
                <svg
                  aria-hidden="true"
                  class="native svg-icon iconGoogle"
                  width="18"
                  height="18"
                  viewBox="0 0 18 18"
                >
                  <path
                    d="M16.51 8H8.98v3h4.3c-.18 1-.74 1.48-1.6 2.04v2.01h2.6a7.8 7.8 0 002.38-5.88c0-.57-.05-.66-.15-1.18z"
                    fill="#4285F4"
                  ></path>
                  <path
                    d="M8.98 17c2.16 0 3.97-.72 5.3-1.94l-2.6-2a4.8 4.8 0 01-7.18-2.54H1.83v2.07A8 8 0 008.98 17z"
                    fill="#34A853"
                  ></path>
                  <path
                    d="M4.5 10.52a4.8 4.8 0 010-3.04V5.41H1.83a8 8 0 000 7.18l2.67-2.07z"
                    fill="#FBBC05"
                  ></path>
                  <path
                    d="M8.98 4.18c1.17 0 2.23.4 3.06 1.2l2.3-2.3A8 8 0 001.83 5.4L4.5 7.49a4.77 4.77 0 014.48-3.3z"
                    fill="#EA4335"
                  ></path>
                </svg>
                Log in with Google{" "}
              </button>
              <button
                class="grid--cell s-btn s-btn__icon s-btn__github bar-md ba bc-black-100"
                data-provider="github"
                data-oauthserver="https://github.com/login/oauth/authorize"
                data-oauthversion="2.0"
              >
                <svg
                  aria-hidden="true"
                  class="svg-icon iconGitHub"
                  width="18"
                  height="18"
                  viewBox="0 0 18 18"
                >
                  <path
                    d="M9 1a8 8 0 00-2.53 15.59c.4.07.55-.17.55-.38l-.01-1.49c-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82a7.42 7.42 0 014 0c1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48l-.01 2.2c0 .21.15.46.55.38A8.01 8.01 0 009 1z"
                    fill="#010101"
                  ></path>
                </svg>
                Log in with GitHub{" "}
              </button>
              <button
                class="grid--cell s-btn s-btn__icon s-btn__facebook bar-md"
                data-provider="facebook"
                data-oauthserver="https://www.facebook.com/v2.0/dialog/oauth"
                data-oauthversion="2.0"
              >
                <svg
                  aria-hidden="true"
                  class="svg-icon iconFacebook"
                  width="18"
                  height="18"
                  viewBox="0 0 18 18"
                >
                  <path
                    d="M3 1a2 2 0 00-2 2v12c0 1.1.9 2 2 2h12a2 2 0 002-2V3a2 2 0 00-2-2H3zm6.55 16v-6.2H7.46V8.4h2.09V6.61c0-2.07 1.26-3.2 3.1-3.2.88 0 1.64.07 1.87.1v2.16h-1.29c-1 0-1.19.48-1.19 1.18V8.4h2.39l-.31 2.42h-2.08V17h-2.5z"
                    fill="#4167B2"
                  ></path>
                </svg>
                Log in with Facebook{" "}
              </button>
            </div>

            <div
              id="formContainer"
              class="mx-auto mb24 p24 wmx3 bg-white bar-lg auth-shadow mb24"
            >
              <form
                id="login-form"
                class="grid fd-column gs12 gsy"
                action=""
                method="POST"
              >
                <div class="grid fd-column gs4 gsy js-auth-item ">
                  <label class="grid--cell s-label" for="email">
                    Email
                  </label>
                  <div class="grid ps-relative">
                    <input
                      class="s-input"
                      id="email"
                      type="email"
                      size="30"
                      maxlength="100"
                      name="email"
                    />
                  </div>
                  <p class="grid--cell s-input-message js-error-message d-none"></p>
                </div>
                <div class="grid fd-column-reverse gs4 gsy js-auth-item ">
                  <p class="grid--cell s-input-message js-error-message d-none"></p>

                  <div class="grid ps-relative js-password">
                    <input
                      class="grid--cell s-input"
                      type="password"
                      autocomplete="off"
                      name="password"
                      id="password"
                    />
                  </div>
                  <div class="grid ai-center ps-relative jc-space-between">
                    <label class="grid--cell s-label" for="password">
                      Password
                    </label>

                    <a
                      class="grid--cell s-link fs-caption"
                      href="/users/account-recovery"
                    >
                      Forgot password?
                    </a>
                  </div>
                </div>
                <div class="grid gs4 gsy fd-column js-auth-item ">
                  <button
                    class="grid--cell s-btn s-btn__primary"
                    id="submit-button"
                    name="submit-button"
                    style={{backgroundColor: 'green', marginTop: '12px'}}
                  >
                    Log in
                  </button>
                  <p class="grid--cell s-input-message js-error-message d-none"></p>
                </div>

                <input type="hidden" id="oauth_version" name="oauth_version" />
                <input type="hidden" id="oauth_server" name="oauth_server" />
              </form>
            </div>

            <div class="mx-auto ta-center fs-body1 p16 pb0 mb24 w100 wmx3">
              Don’t have an account?{" "}
              <a href="/users/signup?ssrc=head&amp;returnurl=https%3a%2f%2fstackoverflow.com%2fquestions%2f29973357%2fhow-do-you-format-code-in-visual-studio-code-vscode">
                Sign up
              </a>
            </div>
          </div>
        </div>
      </div>
      </html>
    );
  }
}

export default Login;
