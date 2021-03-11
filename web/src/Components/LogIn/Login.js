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
