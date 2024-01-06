# Account Subsystem Design

## Technical Analysis

### third-party login plathform

Third-party login platforms need to integrate with APIs from various third-party service providers. Different service providers may use different API protocols, data formats, and versions, making it a challenge to ensure compatibility between the platform and different third-party service providers' APIs.

---

Considering the integration of login methods from multiple third-party platforms, we have researched the following login solutions:


|| feature | platform | docs |
|:----:|----|----|:----:|
|Goth (4.6k star) | 1. Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications. <br> 2. Unlike other similar packages, Goth, lets you write OAuth, OAuth2, or any other protocol providers, as long as they implement the Provider and Session interfaces. | [✓]Facebook  <br>  [✓]Twitter  <br>  [✗]Wechat  <br>  [✓]WeCom(but) | [github.com/markbates/goth](https://github.com/markbates/goth) |
| Casdoor (7.7k star) | A UI-first Identity Access Management (IAM) / Single-Sign-On (SSO) platform supporting OAuth 2.0, OIDC, SAML and CAS, integrated with Casbin RBAC and ABAC permission management | [✓]Facebook  <br>  [✓]Twitter  <br>  [✓]Wechat | [casdoor.org](https://casdoor.org)|
| Ory Hydra (14.7k star) | Highly configurable, suitable for large and complex security requirements |  [✓]Facebook  <br>  [✓]Twitter  <br>  [✗]Wechat | [github.com/ory/hydra](https://github.com/ory/hydra) |
#### Platform analysis 
##### Twitter:X

[Twiter Login Docs](https://developer.twitter.com/en/docs/authentication/guides/log-in-with-twitter)

> The platform has a subscription with multiple development privileges

1. Free(0): **1 app Id**, Login with X ...
2. Basic($100 per month): **2 app Ids**, Login with X ...
3. Pro($5000 per month): **3 app Ids**, Login with X ...
4. Enterprise(...): **all**

##### Wechat

[WeChat Login Docs](https://developers.weixin.qq.com/doc/oplatform/en/Mobile_App/WeChat_Login/Development_Guide.html)

Register as a developer account on the WeChat Open Platform, obtain an approved website application with the corresponding AppID and AppSecret, apply for WeChat Login, and once approved, you can proceed with the integration process.

##### Facebook

[Facebook Login Docs](https://developers.facebook.com/docs/facebook-login/overview)

Create a Facebook developer account and integrate it with our website to obtain the credentials without any restrictions or limitations.

#### result

After researching goth, cashdoor, and Ory Hydra, the following conclusions have been drawn:

1. goth provides example function calls for each platform, allowing integration with just a few lines of code. However, account functionality needs to be implemented separately.
2. Ory Hydra is complex and requires high configuration, but it does not support platforms like WeChat or WeChat for Enterprise as they primarily focus on international platforms.
3. cashdoor offers simple deployment, configuration, and integration, and it provides the necessary functionalities for a user system. We do not need to develop any additional features related to account verification.

Based on the above findings, it has been decided to use **Cashdoor** for quick implementation of account login and third-party platform login. 
