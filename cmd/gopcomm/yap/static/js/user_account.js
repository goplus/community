const authInfo = {
    Google: {
      scope: "profile+email",
      endpoint: "https://accounts.google.com/signin/oauth",
    },
    GitHub: {
      scope: "user:email+read:user",
      endpoint: "https://github.com/login/oauth/authorize",
    },
    QQ: {
      scope: "get_user_info",
      endpoint: "https://graph.qq.com/oauth2.0/authorize",
    },
    WeChat: {
      scope: "snsapi_login",
      endpoint: "https://open.weixin.qq.com/connect/qrconnect",
      mpScope: "snsapi_userinfo",
      mpEndpoint: "https://open.weixin.qq.com/connect/oauth2/authorize",
    },
    WeChatMiniProgram: {
      endpoint: "https://mp.weixin.qq.com/",
    },
    Facebook: {
      scope: "email,public_profile",
      endpoint: "https://www.facebook.com/dialog/oauth",
    },
    DingTalk: {
      scope: "openid",
      endpoint: "https://login.dingtalk.com/oauth2/auth",
    },
    Weibo: {
      scope: "email",
      endpoint: "https://api.weibo.com/oauth2/authorize",
    },
    Gitee: {
      scope: "user_info%20emails",
      endpoint: "https://gitee.com/oauth/authorize",
    },
    LinkedIn: {
      scope: "r_liteprofile%20r_emailaddress",
      endpoint: "https://www.linkedin.com/oauth/v2/authorization",
    },
    WeCom: {
      scope: "snsapi_userinfo",
      endpoint: "https://open.work.weixin.qq.com/wwopen/sso/3rd_qrConnect",
      silentEndpoint: "https://open.weixin.qq.com/connect/oauth2/authorize",
      internalEndpoint: "https://open.work.weixin.qq.com/wwopen/sso/qrConnect",
    },
    Lark: {
      // scope: "email",
      endpoint: "https://open.feishu.cn/open-apis/authen/v1/index",
    },
    GitLab: {
      scope: "read_user+profile",
      endpoint: "https://gitlab.com/oauth/authorize",
    },
    ADFS: {
      scope: "openid",
      endpoint: "http://example.com",
    },
    Baidu: {
      scope: "basic",
      endpoint: "http://openapi.baidu.com/oauth/2.0/authorize",
    },
    Alipay: {
      scope: "basic",
      endpoint: "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm",
    },
    Casdoor: {
      scope: "openid%20profile%20email",
      endpoint: "http://example.com",
    },
    Infoflow: {
      endpoint: "https://xpc.im.baidu.com/oauth2/authorize",
    },
    Apple: {
      scope: "name%20email",
      endpoint: "https://appleid.apple.com/auth/authorize",
    },
    AzureAD: {
      scope: "user.read",
      endpoint: "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
    },
    AzureADB2C: {
      scope: "openid",
      endpoint: "https://tenant.b2clogin.com/tenant.onmicrosoft.com/userflow/oauth2/v2.0/authorize",
    },
    Slack: {
      scope: "users:read",
      endpoint: "https://slack.com/oauth/authorize",
    },
    Steam: {
      endpoint: "https://steamcommunity.com/openid/login",
    },
    Okta: {
      scope: "openid%20profile%20email",
      endpoint: "http://example.com",
    },
    Douyin: {
      scope: "user_info",
      endpoint: "https://open.douyin.com/platform/oauth/connect",
    },
    Custom: {
      endpoint: "https://example.com/",
    },
    Bilibili: {
      endpoint: "https://passport.bilibili.com/register/pc_oauth2.html",
    },
    Line: {
      scope: "profile%20openid%20email",
      endpoint: "https://access.line.me/oauth2/v2.1/authorize",
    },
    Amazon: {
      scope: "profile",
      endpoint: "https://www.amazon.com/ap/oa",
    },
    Auth0: {
      scope: "openid%20profile%20email",
      endpoint: "http://auth0.com/authorize",
    },
    BattleNet: {
      scope: "openid",
      endpoint: "https://oauth.battlenet.com.cn/authorize",
    },
    Bitbucket: {
      scope: "account",
      endpoint: "https://bitbucket.org/site/oauth2/authorize",
    },
    Box: {
      scope: "root_readwrite",
      endpoint: "https://account.box.com/api/oauth2/authorize",
    },
    CloudFoundry: {
      scope: "cloud_controller.read",
      endpoint: "https://login.cloudfoundry.org/oauth/authorize",
    },
    Dailymotion: {
      scope: "userinfo",
      endpoint: "https://api.dailymotion.com/oauth/authorize",
    },
    Deezer: {
      scope: "basic_access",
      endpoint: "https://connect.deezer.com/oauth/auth.php",
    },
    DigitalOcean: {
      scope: "read",
      endpoint: "https://cloud.digitalocean.com/v1/oauth/authorize",
    },
    Discord: {
      scope: "identify%20email",
      endpoint: "https://discord.com/api/oauth2/authorize",
    },
    Dropbox: {
      scope: "account_info.read",
      endpoint: "https://www.dropbox.com/oauth2/authorize",
    },
    EveOnline: {
      scope: "publicData",
      endpoint: "https://login.eveonline.com/oauth/authorize",
    },
    Fitbit: {
      scope: "activity%20heartrate%20location%20nutrition%20profile%20settings%20sleep%20social%20weight",
      endpoint: "https://www.fitbit.com/oauth2/authorize",
    },
    Gitea: {
      scope: "user:email",
      endpoint: "https://gitea.com/login/oauth/authorize",
    },
    Heroku: {
      scope: "global",
      endpoint: "https://id.heroku.com/oauth/authorize",
    },
    InfluxCloud: {
      scope: "read:org",
      endpoint: "https://cloud2.influxdata.com/oauth/authorize",
    },
    Instagram: {
      scope: "user_profile",
      endpoint: "https://api.instagram.com/oauth/authorize",
    },
    Intercom: {
      scope: "user.read",
      endpoint: "https://app.intercom.com/oauth",
    },
    Kakao: {
      scope: "account_email",
      endpoint: "https://kauth.kakao.com/oauth/authorize",
    },
    Lastfm: {
      scope: "user_read",
      endpoint: "https://www.last.fm/api/auth",
    },
    Mailru: {
      scope: "userinfo",
      endpoint: "https://oauth.mail.ru/login",
    },
    Meetup: {
      scope: "basic",
      endpoint: "https://secure.meetup.com/oauth2/authorize",
    },
    MicrosoftOnline: {
      scope: "openid%20profile%20email",
      endpoint: "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
    },
    Naver: {
      scope: "profile",
      endpoint: "https://nid.naver.com/oauth2.0/authorize",
    },
    Nextcloud: {
      scope: "openid%20profile%20email",
      endpoint: "https://cloud.example.org/apps/oauth2/authorize",
    },
    OneDrive: {
      scope: "offline_access%20onedrive.readonly",
      endpoint: "https://login.live.com/oauth20_authorize.srf",
    },
    Oura: {
      scope: "personal",
      endpoint: "https://cloud.ouraring.com/oauth/authorize",
    },
    Patreon: {
      scope: "identity",
      endpoint: "https://www.patreon.com/oauth2/authorize",
    },
    PayPal: {
      scope: "openid%20profile%20email",
      endpoint: "https://www.sandbox.paypal.com/connect",
    },
    SalesForce: {
      scope: "openid%20profile%20email",
      endpoint: "https://login.salesforce.com/services/oauth2/authorize",
    },
    Shopify: {
      scope: "read_products",
      endpoint: "https://myshopify.com/admin/oauth/authorize",
    },
    Soundcloud: {
      scope: "non-expiring",
      endpoint: "https://api.soundcloud.com/connect",
    },
    Spotify: {
      scope: "user-read-email",
      endpoint: "https://accounts.spotify.com/authorize",
    },
    Strava: {
      scope: "read",
      endpoint: "https://www.strava.com/oauth/authorize",
    },
    Stripe: {
      scope: "read_only",
      endpoint: "https://connect.stripe.com/oauth/authorize",
    },
    TikTok: {
      scope: "user.info.basic",
      endpoint: "https://www.tiktok.com/auth/authorize/",
    },
    Tumblr: {
      scope: "email",
      endpoint: "https://www.tumblr.com/oauth2/authorize",
    },
    Twitch: {
      scope: "user_read",
      endpoint: "https://id.twitch.tv/oauth2/authorize",
    },
    Twitter: {
      scope: "users.read%20tweet.read",
      endpoint: "https://twitter.com/i/oauth2/authorize",
    },
    Typetalk: {
      scope: "my",
      endpoint: "https://typetalk.com/oauth2/authorize",
    },
    Uber: {
      scope: "profile",
      endpoint: "https://login.uber.com/oauth/v2/authorize",
    },
    VK: {
      scope: "email",
      endpoint: "https://oauth.vk.com/authorize",
    },
    Wepay: {
      scope: "manage_accounts%20view_user",
      endpoint: "https://www.wepay.com/v2/oauth2/authorize",
    },
    Xero: {
      scope: "openid%20profile%20email",
      endpoint: "https://login.xero.com/identity/connect/authorize",
    },
    Yahoo: {
      scope: "openid%20profile%20email",
      endpoint: "https://api.login.yahoo.com/oauth2/request_auth",
    },
    Yammer: {
      scope: "user",
      endpoint: "https://www.yammer.com/oauth2/authorize",
    },
    Yandex: {
      scope: "login:email",
      endpoint: "https://oauth.yandex.com/authorize",
    },
    Zoom: {
      scope: "user:read",
      endpoint: "https://zoom.us/oauth/authorize",
    },
    MetaMask: {
      scope: "",
      endpoint: "",
    },
    Web3Onboard: {
      scope: "",
      endpoint: "",
    },
  };

  const providers = [
    {
        "owner": "admin",
        "name": "provider_github",
        "createdTime": "2024-01-22T11:15:22+08:00",
        "displayName": "New Provider - github",
        "category": "OAuth",
        "type": "GitHub",
        "subType": "",
        "method": "Normal",
        "clientId": "Iv1.de7a592b6cc33131",
        "clientSecret": "***",
        "clientId2": "",
        "clientSecret2": "",
        "cert": "",
        "customAuthUrl": "",
        "customTokenUrl": "",
        "customUserInfoUrl": "",
        "customLogo": "",
        "scopes": "",
        "userMapping": {},
        "host": "",
        "port": 0,
        "disableSsl": false,
        "title": "",
        "content": "",
        "receiver": "",
        "regionId": "",
        "signName": "",
        "templateCode": "",
        "appId": "",
        "endpoint": "",
        "intranetEndpoint": "",
        "domain": "",
        "bucket": "",
        "pathPrefix": "",
        "metadata": "",
        "idP": "",
        "issuerUrl": "",
        "enableSignAuthnRequest": false,
        "providerUrl": "https://github.com/organizations/xxx/settings/applications/1234567"
    },
    {
        "owner": "admin",
        "name": "provider_twitter",
        "createdTime": "2024-01-22T11:20:19+08:00",
        "displayName": "New Provider - twitter",
        "category": "OAuth",
        "type": "Twitter",
        "subType": "",
        "method": "Normal",
        "clientId": "VmNJS2otVTlZYUVLNEVVZ0ZJc2w6MTpjaQ",
        "clientSecret": "***",
        "clientId2": "",
        "clientSecret2": "",
        "cert": "",
        "customAuthUrl": "",
        "customTokenUrl": "",
        "customUserInfoUrl": "",
        "customLogo": "",
        "scopes": "",
        "userMapping": {},
        "host": "",
        "port": 0,
        "disableSsl": false,
        "title": "",
        "content": "",
        "receiver": "",
        "regionId": "",
        "signName": "",
        "templateCode": "",
        "appId": "",
        "endpoint": "",
        "intranetEndpoint": "",
        "domain": "",
        "bucket": "",
        "pathPrefix": "",
        "metadata": "",
        "idP": "",
        "issuerUrl": "",
        "enableSignAuthnRequest": false,
        "providerUrl": "https://github.com/organizations/xxx/settings/applications/1234567"
    },
    {
        "owner": "admin",
        "name": "provider_facebook",
        "createdTime": "2024-01-22T11:19:03+08:00",
        "displayName": "New Provider - facebook",
        "category": "OAuth",
        "type": "Facebook",
        "subType": "",
        "method": "Normal",
        "clientId": "927083362474156",
        "clientSecret": "***",
        "clientId2": "",
        "clientSecret2": "",
        "cert": "",
        "customAuthUrl": "",
        "customTokenUrl": "",
        "customUserInfoUrl": "",
        "customLogo": "",
        "scopes": "",
        "userMapping": {},
        "host": "",
        "port": 0,
        "disableSsl": false,
        "title": "",
        "content": "",
        "receiver": "",
        "regionId": "",
        "signName": "",
        "templateCode": "",
        "appId": "",
        "endpoint": "",
        "intranetEndpoint": "",
        "domain": "",
        "bucket": "",
        "pathPrefix": "",
        "metadata": "",
        "idP": "",
        "issuerUrl": "",
        "enableSignAuthnRequest": false,
        "providerUrl": "https://github.com/organizations/xxx/settings/applications/1234567"
    },
    {
        "owner": "admin",
        "name": "provider_wechat",
        "createdTime": "2024-01-25T10:20:42+08:00",
        "displayName": "wechat",
        "category": "OAuth",
        "type": "WeChat",
        "subType": "",
        "method": "Normal",
        "clientId": "wxf6fad0f92d9e8313",
        "clientSecret": "***",
        "clientId2": "wxf6fad0f92d9e8313",
        "clientSecret2": "***",
        "cert": "",
        "customAuthUrl": "",
        "customTokenUrl": "",
        "customUserInfoUrl": "",
        "customLogo": "",
        "scopes": "",
        "userMapping": {},
        "host": "",
        "port": 0,
        "disableSsl": true,
        "title": "",
        "content": "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAAB+ElEQVR42uyYPfLrIAzE10NByRE4Cjfzx804CkegpGC8byQ8SZyX6l9ZHquK418jI1Yr4YknnvhbLCTZAFbMAMiy9CD/8V4AgKkhsWJiTY5l6iEfL+wAjtxazI5laZrmjJAlW4MA+8i2ezksq4AeDytcA3BHQGsyZlZ5Sh1xx4+ivTqg+hBzD2X2NbkWt18CcnFgRCQUGDfrR1wckJLjULkZgR2Rmma5FwDWtHrS1ShHKDXpmKf3zTIBLK4eNyuyhexaXF3F7rnaApjhC0iuCHQNU9fOBFtAxSK9KBTR6nFgSWzPvQAErohkjVsLlE8i8v7p5WwAeRH/IGlSzJBYOiCulgA4cvcFesEAeK5DH2AKQJDHpJUXCD23tPu3DN4GmD2lD+3aebkjcOJLBm0AfJ2dyIQOQfgUcxPA0rUQs9o3TXNTMX/rgwnA1bSRuYvY6RA0j0u23A7YMfyDDEGjJZ3mbguA6MMoxvmV5rfaXx9wYqjFjeoQJPrgaprOLsgCkIazEcHTJYJAJ8t6A+Bj1hP/wKaScmpJBoBjeZi6/pTxgT3kz+9gAjh2IB3qSDv+S9MEcOyjBBgqt43EbwloTbah4pS7ViwC+chP5m4Efq16rg+MrbtqhKTpxcuJr7YFHMtDnUKrcjJW80tArANPPPHEOf4FAAD//zp6rk3/kINSAAAAAElFTkSuQmCC",
        "receiver": "",
        "regionId": "",
        "signName": "",
        "templateCode": "",
        "appId": "",
        "endpoint": "",
        "intranetEndpoint": "",
        "domain": "",
        "bucket": "",
        "pathPrefix": "",
        "metadata": "",
        "idP": "",
        "issuerUrl": "",
        "enableSignAuthnRequest": false,
        "providerUrl": "https://open.weixin.qq.com/connect/oauth2/authorize"
    }
  ]
function getStateFromQueryParams(applicationName, providerName, method, isShortState) {
    let query = window.location.search;
    query = `${query}&application=${encodeURIComponent(applicationName)}&provider=${encodeURIComponent(providerName)}&method=${method}`;
    if (method === "link") {
      query = `${query}&from=${window.location.pathname}`;
    }
  
    if (!isShortState) {
      return btoa(query);
    } else {
      const state = providerName;
      sessionStorage.setItem(state, query);
      return state;
    }
  }
function getAuthUrl(provider, method, res) {
    application = {
        "owner": "admin",
        "name": "application_x8aevk",
        "createdTime": "2024-01-22T11:12:21+08:00",
        "displayName": "New Application - x8aevk",
        "logo": "https://img0.baidu.com/it/u=1037925000,2019845499&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500",
        "homepageUrl": "",
        "description": "",
        "organization": "built-in",
        "cert": "cert-built-in",
        "enablePassword": true,
        "enableSignUp": true,
        "enableSigninSession": false,
        "enableAutoSignin": false,
        "enableCodeSignin": false,
        "enableSamlCompress": false,
        "enableSamlC14n10": false,
        "enableSamlPostBinding": false,
        "enableWebAuthn": false,
        "enableLinkWithEmail": true,
        "orgChoiceMode": "",
        "samlReplyUrl": "",
        "providers": [
            {
                "owner": "",
                "name": "provider_github",
                "canSignUp": true,
                "canSignIn": true,
                "canUnlink": true,
                "prompted": false,
                "signupGroup": "",
                "rule": "None",
                "provider": {
                    "owner": "admin",
                    "name": "provider_github",
                    "createdTime": "2024-01-22T11:15:22+08:00",
                    "displayName": "New Provider - github",
                    "category": "OAuth",
                    "type": "GitHub",
                    "subType": "",
                    "method": "Normal",
                    "clientId": "Iv1.de7a592b6cc33131",
                    "clientSecret": "***",
                    "clientId2": "",
                    "clientSecret2": "",
                    "cert": "",
                    "customAuthUrl": "",
                    "customTokenUrl": "",
                    "customUserInfoUrl": "",
                    "customLogo": "",
                    "scopes": "",
                    "userMapping": {},
                    "host": "",
                    "port": 0,
                    "disableSsl": false,
                    "title": "",
                    "content": "",
                    "receiver": "",
                    "regionId": "",
                    "signName": "",
                    "templateCode": "",
                    "appId": "",
                    "endpoint": "",
                    "intranetEndpoint": "",
                    "domain": "",
                    "bucket": "",
                    "pathPrefix": "",
                    "metadata": "",
                    "idP": "",
                    "issuerUrl": "",
                    "enableSignAuthnRequest": false,
                    "providerUrl": "https://github.com/organizations/xxx/settings/applications/1234567"
                }
            },
            {
                "owner": "",
                "name": "provider_twitter",
                "canSignUp": true,
                "canSignIn": true,
                "canUnlink": true,
                "prompted": false,
                "signupGroup": "",
                "rule": "None",
                "provider": {
                    "owner": "admin",
                    "name": "provider_twitter",
                    "createdTime": "2024-01-22T11:20:19+08:00",
                    "displayName": "New Provider - twitter",
                    "category": "OAuth",
                    "type": "Twitter",
                    "subType": "",
                    "method": "Normal",
                    "clientId": "VmNJS2otVTlZYUVLNEVVZ0ZJc2w6MTpjaQ",
                    "clientSecret": "***",
                    "clientId2": "",
                    "clientSecret2": "",
                    "cert": "",
                    "customAuthUrl": "",
                    "customTokenUrl": "",
                    "customUserInfoUrl": "",
                    "customLogo": "",
                    "scopes": "",
                    "userMapping": {},
                    "host": "",
                    "port": 0,
                    "disableSsl": false,
                    "title": "",
                    "content": "",
                    "receiver": "",
                    "regionId": "",
                    "signName": "",
                    "templateCode": "",
                    "appId": "",
                    "endpoint": "",
                    "intranetEndpoint": "",
                    "domain": "",
                    "bucket": "",
                    "pathPrefix": "",
                    "metadata": "",
                    "idP": "",
                    "issuerUrl": "",
                    "enableSignAuthnRequest": false,
                    "providerUrl": "https://github.com/organizations/xxx/settings/applications/1234567"
                }
            },
            {
                "owner": "",
                "name": "provider_facebook",
                "canSignUp": true,
                "canSignIn": true,
                "canUnlink": true,
                "prompted": false,
                "signupGroup": "",
                "rule": "None",
                "provider": {
                    "owner": "admin",
                    "name": "provider_facebook",
                    "createdTime": "2024-01-22T11:19:03+08:00",
                    "displayName": "New Provider - facebook",
                    "category": "OAuth",
                    "type": "Facebook",
                    "subType": "",
                    "method": "Normal",
                    "clientId": "927083362474156",
                    "clientSecret": "***",
                    "clientId2": "",
                    "clientSecret2": "",
                    "cert": "",
                    "customAuthUrl": "",
                    "customTokenUrl": "",
                    "customUserInfoUrl": "",
                    "customLogo": "",
                    "scopes": "",
                    "userMapping": {},
                    "host": "",
                    "port": 0,
                    "disableSsl": false,
                    "title": "",
                    "content": "",
                    "receiver": "",
                    "regionId": "",
                    "signName": "",
                    "templateCode": "",
                    "appId": "",
                    "endpoint": "",
                    "intranetEndpoint": "",
                    "domain": "",
                    "bucket": "",
                    "pathPrefix": "",
                    "metadata": "",
                    "idP": "",
                    "issuerUrl": "",
                    "enableSignAuthnRequest": false,
                    "providerUrl": "https://github.com/organizations/xxx/settings/applications/1234567"
                }
            },
            {
                "owner": "",
                "name": "provider_wechat",
                "canSignUp": true,
                "canSignIn": true,
                "canUnlink": true,
                "prompted": false,
                "signupGroup": "",
                "rule": "None",
                "provider": {
                    "owner": "admin",
                    "name": "provider_wechat",
                    "createdTime": "2024-01-25T10:20:42+08:00",
                    "displayName": "wechat",
                    "category": "OAuth",
                    "type": "WeChat",
                    "subType": "",
                    "method": "Normal",
                    "clientId": "wxf6fad0f92d9e8313",
                    "clientSecret": "***",
                    "clientId2": "wxf6fad0f92d9e8313",
                    "clientSecret2": "***",
                    "cert": "",
                    "customAuthUrl": "",
                    "customTokenUrl": "",
                    "customUserInfoUrl": "",
                    "customLogo": "",
                    "scopes": "",
                    "userMapping": {},
                    "host": "",
                    "port": 0,
                    "disableSsl": true,
                    "title": "",
                    "content": "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAAB+ElEQVR42uyYPfLrIAzE10NByRE4Cjfzx804CkegpGC8byQ8SZyX6l9ZHquK418jI1Yr4YknnvhbLCTZAFbMAMiy9CD/8V4AgKkhsWJiTY5l6iEfL+wAjtxazI5laZrmjJAlW4MA+8i2ezksq4AeDytcA3BHQGsyZlZ5Sh1xx4+ivTqg+hBzD2X2NbkWt18CcnFgRCQUGDfrR1wckJLjULkZgR2Rmma5FwDWtHrS1ShHKDXpmKf3zTIBLK4eNyuyhexaXF3F7rnaApjhC0iuCHQNU9fOBFtAxSK9KBTR6nFgSWzPvQAErohkjVsLlE8i8v7p5WwAeRH/IGlSzJBYOiCulgA4cvcFesEAeK5DH2AKQJDHpJUXCD23tPu3DN4GmD2lD+3aebkjcOJLBm0AfJ2dyIQOQfgUcxPA0rUQs9o3TXNTMX/rgwnA1bSRuYvY6RA0j0u23A7YMfyDDEGjJZ3mbguA6MMoxvmV5rfaXx9wYqjFjeoQJPrgaprOLsgCkIazEcHTJYJAJ8t6A+Bj1hP/wKaScmpJBoBjeZi6/pTxgT3kz+9gAjh2IB3qSDv+S9MEcOyjBBgqt43EbwloTbah4pS7ViwC+chP5m4Efq16rg+MrbtqhKTpxcuJr7YFHMtDnUKrcjJW80tArANPPPHEOf4FAAD//zp6rk3/kINSAAAAAElFTkSuQmCC",
                    "receiver": "",
                    "regionId": "",
                    "signName": "",
                    "templateCode": "",
                    "appId": "",
                    "endpoint": "",
                    "intranetEndpoint": "",
                    "domain": "",
                    "bucket": "",
                    "pathPrefix": "",
                    "metadata": "",
                    "idP": "",
                    "issuerUrl": "",
                    "enableSignAuthnRequest": false,
                    "providerUrl": "https://open.weixin.qq.com/connect/oauth2/authorize"
                }
            }
        ],
        "signinMethods": [
            {
                "name": "Password",
                "displayName": "Password",
                "rule": "All"
            }
        ],
        "signupItems": [
            {
                "name": "ID",
                "visible": false,
                "required": true,
                "prompted": false,
                "label": "",
                "placeholder": "",
                "regex": "",
                "rule": "Random"
            },
            {
                "name": "Username",
                "visible": true,
                "required": true,
                "prompted": false,
                "label": "",
                "placeholder": "",
                "regex": "",
                "rule": "None"
            },
            {
                "name": "Display name",
                "visible": true,
                "required": true,
                "prompted": false,
                "label": "",
                "placeholder": "",
                "regex": "",
                "rule": "None"
            },
            {
                "name": "Password",
                "visible": true,
                "required": true,
                "prompted": false,
                "label": "",
                "placeholder": "",
                "regex": "",
                "rule": "None"
            },
            {
                "name": "Confirm password",
                "visible": true,
                "required": true,
                "prompted": false,
                "label": "",
                "placeholder": "",
                "regex": "",
                "rule": "None"
            },
            {
                "name": "Agreement",
                "visible": true,
                "required": true,
                "prompted": false,
                "label": "",
                "placeholder": "",
                "regex": "",
                "rule": "None"
            }
        ],
        "grantTypes": [
            "authorization_code"
        ],
        "organizationObj": {
            "owner": "admin",
            "name": "built-in",
            "createdTime": "2024-01-22T01:27:18Z",
            "displayName": "Built-in Organization",
            "websiteUrl": "https://example.com",
            "favicon": "https://cdn.casbin.org/img/casbin/favicon.ico",
            "passwordType": "plain",
            "passwordSalt": "",
            "passwordOptions": [
                "AtLeast6"
            ],
            "countryCodes": [
                "US",
                "ES",
                "FR",
                "DE",
                "GB",
                "CN",
                "JP",
                "KR",
                "VN",
                "ID",
                "SG",
                "IN"
            ],
            "defaultAvatar": "https://cdn.casbin.org/img/casbin.svg",
            "defaultApplication": "",
            "tags": [],
            "languages": [
                "en",
                "zh",
                "es",
                "fr",
                "de",
                "id",
                "ja",
                "ko",
                "ru",
                "vi",
                "pt"
            ],
            "themeData": null,
            "masterPassword": "",
            "defaultPassword": "",
            "masterVerificationCode": "",
            "initScore": 2000,
            "enableSoftDeletion": false,
            "isProfilePublic": false,
            "mfaItems": null,
            "accountItems": [
                {
                    "name": "Organization",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Admin"
                },
                {
                    "name": "ID",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Immutable"
                },
                {
                    "name": "Name",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Admin"
                },
                {
                    "name": "Display name",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Self"
                },
                {
                    "name": "Avatar",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Self"
                },
                {
                    "name": "User type",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Admin"
                },
                {
                    "name": "Password",
                    "visible": true,
                    "viewRule": "Self",
                    "modifyRule": "Self"
                },
                {
                    "name": "Email",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Self"
                },
                {
                    "name": "Phone",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Self"
                },
                {
                    "name": "Country code",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Admin"
                },
                {
                    "name": "Country/Region",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Self"
                },
                {
                    "name": "Location",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Self"
                },
                {
                    "name": "Affiliation",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Self"
                },
                {
                    "name": "Title",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Self"
                },
                {
                    "name": "Homepage",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Self"
                },
                {
                    "name": "Bio",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Self"
                },
                {
                    "name": "Tag",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Admin"
                },
                {
                    "name": "Signup application",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Admin"
                },
                {
                    "name": "Roles",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Immutable"
                },
                {
                    "name": "Permissions",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Immutable"
                },
                {
                    "name": "Groups",
                    "visible": true,
                    "viewRule": "Public",
                    "modifyRule": "Admin"
                },
                {
                    "name": "3rd-party logins",
                    "visible": true,
                    "viewRule": "Self",
                    "modifyRule": "Self"
                },
                {
                    "name": "Properties",
                    "visible": false,
                    "viewRule": "Admin",
                    "modifyRule": "Admin"
                },
                {
                    "name": "Is admin",
                    "visible": true,
                    "viewRule": "Admin",
                    "modifyRule": "Admin"
                },
                {
                    "name": "Is forbidden",
                    "visible": true,
                    "viewRule": "Admin",
                    "modifyRule": "Admin"
                },
                {
                    "name": "Is deleted",
                    "visible": true,
                    "viewRule": "Admin",
                    "modifyRule": "Admin"
                },
                {
                    "name": "Multi-factor authentication",
                    "visible": true,
                    "viewRule": "Self",
                    "modifyRule": "Self"
                },
                {
                    "name": "WebAuthn credentials",
                    "visible": true,
                    "viewRule": "Self",
                    "modifyRule": "Self"
                },
                {
                    "name": "Managed accounts",
                    "visible": true,
                    "viewRule": "Self",
                    "modifyRule": "Self"
                }
            ]
        },
        "certPublicKey": "",
        "tags": [],
        "samlAttributes": [],
        "clientId": "389313df51ffd2093b2f",
        "clientSecret": "bd22db46e4ab52ab0d0ac9cbd771a21775039409",
        "redirectUris": [
            "http://localhost:8080/callback",
            "http://goplus-community-dev.community.svc.jfcs-qa1.local:8080/",
            "http://goplus-community-asklv.community.svc.jfcs-qa1.local:8080/",
            "http://10.212.223.75:8080//callback",
            "http://10.212.196.141:8080//callback"
        ],
        "tokenFormat": "JWT",
        "tokenFields": [],
        "expireInHours": 168,
        "refreshExpireInHours": 168,
        "signupUrl": "",
        "signinUrl": "",
        "forgetUrl": "",
        "affiliationUrl": "",
        "termsOfUse": "",
        "signupHtml": "",
        "signinHtml": "",
        "themeData": {
            "themeType": "default",
            "colorPrimary": "#1677FF",
            "borderRadius": 3,
            "isCompact": false,
            "isEnabled": true
        },
        "formCss": "<style>\n  .login-panel{\n    padding: 40px 70px 0 70px;\n    border-radius: 10px;\n    background-color: #ffffff;\n    box-shadow: 0 0 30px 20px rgba(0, 0, 0, 0.20);\n}\n\n\n</style>",
        "formCssMobile": "",
        "formOffset": 2,
        "formSideHtml": "",
        "formBackgroundUrl": "",
        "failedSigninLimit": 5,
        "failedSigninFrozenTime": 15
    }
    if (provider === null) {
      return "";
    }
    let endpoint = authInfo[provider.type].endpoint;
    let redirectUri = `https://casdoor-community.qiniu.io/callback`;
    const scope = authInfo[provider.type].scope;
  
    const isShortState = (provider.type === "WeChat" && navigator.userAgent.includes("MicroMessenger")) || (provider.type === "Twitter");
    const state = getStateFromQueryParams(application.name, provider.name, method, isShortState);
    const codeChallenge = "P3S-a7dr8bgM4bF6vOyiKkKETDl16rcAzao9F8UIL1Y"; // SHA256(Base64-URL-encode("casdoor-verifier"))
  
    if (provider.type === "AzureAD") {
      if (provider.domain !== "") {
        endpoint = endpoint.replace("common", provider.domain);
      }
    } else if (provider.type === "Apple") {
      redirectUri = `https://casdoor-community.qiniu.io/api/callback`;
    }
  
    if (provider.type === "Google" || provider.type === "GitHub" || provider.type === "QQ" || provider.type === "Facebook"
      || provider.type === "Weibo" || provider.type === "Gitee" || provider.type === "LinkedIn" || provider.type === "GitLab" || provider.type === "AzureAD"
      || provider.type === "Slack" || provider.type === "Line" || provider.type === "Amazon" || provider.type === "Auth0" || provider.type === "BattleNet"
      || provider.type === "Bitbucket" || provider.type === "Box" || provider.type === "CloudFoundry" || provider.type === "Dailymotion"
      || provider.type === "DigitalOcean" || provider.type === "Discord" || provider.type === "Dropbox" || provider.type === "EveOnline" || provider.type === "Gitea"
      || provider.type === "Heroku" || provider.type === "InfluxCloud" || provider.type === "Instagram" || provider.type === "Intercom" || provider.type === "Kakao"
      || provider.type === "MailRu" || provider.type === "Meetup" || provider.type === "MicrosoftOnline" || provider.type === "Naver" || provider.type === "Nextcloud"
      || provider.type === "OneDrive" || provider.type === "Oura" || provider.type === "Patreon" || provider.type === "PayPal" || provider.type === "SalesForce"
      || provider.type === "SoundCloud" || provider.type === "Spotify" || provider.type === "Strava" || provider.type === "Stripe" || provider.type === "Tumblr"
      || provider.type === "Twitch" || provider.type === "Typetalk" || provider.type === "Uber" || provider.type === "VK" || provider.type === "Wepay"
      || provider.type === "Xero" || provider.type === "Yahoo" || provider.type === "Yammer" || provider.type === "Yandex" || provider.type === "Zoom") {
      return `${endpoint}?client_id=${provider.clientId}&redirect_uri=${redirectUri}&scope=${scope}&response_type=code&state=${state}`;
    } else if (provider.type === "AzureADB2C") {
      return `https://${provider.domain}.b2clogin.com/${provider.domain}.onmicrosoft.com/${provider.appId}/oauth2/v2.0/authorize?client_id=${provider.clientId}&nonce=defaultNonce&redirect_uri=${encodeURIComponent(redirectUri)}&scope=${scope}&response_type=code&state=${state}&prompt=login`;
    } else if (provider.type === "DingTalk") {
      return `${endpoint}?client_id=${provider.clientId}&redirect_uri=${redirectUri}&scope=${scope}&response_type=code&prompt=consent&state=${state}`;
    } else if (provider.type === "WeChat") {
      if (provider.clientId !== "") {
        return `${endpoint}?appid=${provider.clientId}&redirect_uri=${redirectUri}&scope=${scope}&response_type=code&state=${state}#wechat_redirect`;
      }
      return `/callback?appid=${provider.clientId2}&redirect_uri=${redirectUri}&scope=${scope}&response_type=code&state=${state}&code=${res?.data2 || ""}`;
    } else if (provider.type === "WeCom") {
      if (provider.subType === "Internal") {
        if (provider.method === "Silent") {
          endpoint = authInfo[provider.type].silentEndpoint;
          return `${endpoint}?appid=${provider.clientId}&redirect_uri=${redirectUri}&state=${state}&scope=${scope}&response_type=code#wechat_redirect`;
        } else if (provider.method === "Normal") {
          endpoint = authInfo[provider.type].internalEndpoint;
          return `${endpoint}?appid=${provider.clientId}&agentid=${provider.appId}&redirect_uri=${redirectUri}&state=${state}&usertype=member`;
        } else {
          return `https://error:not-supported-provider-method:${provider.method}`;
        }
      } else if (provider.subType === "Third-party") {
        if (provider.method === "Silent") {
          endpoint = authInfo[provider.type].silentEndpoint;
          return `${endpoint}?appid=${provider.clientId}&redirect_uri=${redirectUri}&state=${state}&scope=${scope}&response_type=code#wechat_redirect`;
        } else if (provider.method === "Normal") {
          return `${endpoint}?appid=${provider.clientId}&redirect_uri=${redirectUri}&state=${state}&usertype=member`;
        } else {
          return `https://error:not-supported-provider-method:${provider.method}`;
        }
      } else {
        return `https://error:not-supported-provider-sub-type:${provider.subType}`;
      }
    } else if (provider.type === "Lark") {
      return `${endpoint}?app_id=${provider.clientId}&redirect_uri=${redirectUri}&state=${state}`;
    } else if (provider.type === "ADFS") {
      return `${provider.domain}/adfs/oauth2/authorize?client_id=${provider.clientId}&redirect_uri=${redirectUri}&state=${state}&response_type=code&nonce=casdoor&scope=openid`;
    } else if (provider.type === "Baidu") {
      return `${endpoint}?client_id=${provider.clientId}&redirect_uri=${redirectUri}&state=${state}&response_type=code&scope=${scope}&display=popup`;
    } else if (provider.type === "Alipay") {
      return `${endpoint}?app_id=${provider.clientId}&scope=auth_user&redirect_uri=${redirectUri}&state=${state}&response_type=code&scope=${scope}&display=popup`;
    } else if (provider.type === "Casdoor") {
      return `${provider.domain}/login/oauth/authorize?client_id=${provider.clientId}&redirect_uri=${redirectUri}&state=${state}&response_type=code&scope=${scope}`;
    } else if (provider.type === "Infoflow") {
      return `${endpoint}?appid=${provider.clientId}&redirect_uri=${redirectUri}?state=${state}`;
    } else if (provider.type === "Apple") {
      return `${endpoint}?client_id=${provider.clientId}&redirect_uri=${redirectUri}&state=${state}&response_type=code%20id_token&scope=${scope}&response_mode=form_post`;
    } else if (provider.type === "Steam") {
      return `${endpoint}?openid.claimed_id=http://specs.openid.net/auth/2.0/identifier_select&openid.identity=http://specs.openid.net/auth/2.0/identifier_select&openid.mode=checkid_setup&openid.ns=http://specs.openid.net/auth/2.0&openid.realm=https://casdoor-community.qiniu.io&openid.return_to=${redirectUri}?state=${state}`;
    } else if (provider.type === "Okta") {
      return `${provider.domain}/v1/authorize?client_id=${provider.clientId}&redirect_uri=${redirectUri}&state=${state}&response_type=code&scope=${scope}`;
    } else if (provider.type === "Douyin" || provider.type === "TikTok") {
      return `${endpoint}?client_key=${provider.clientId}&redirect_uri=${redirectUri}&state=${state}&response_type=code&scope=${scope}`;
    } else if (provider.type === "Custom") {
      return `${provider.customAuthUrl}?client_id=${provider.clientId}&redirect_uri=${redirectUri}&scope=${provider.scopes}&response_type=code&state=${state}`;
    } else if (provider.type === "Bilibili") {
      return `${endpoint}#/?client_id=${provider.clientId}&return_url=${redirectUri}&state=${state}&response_type=code`;
    } else if (provider.type === "Deezer") {
      return `${endpoint}?app_id=${provider.clientId}&redirect_uri=${redirectUri}&perms=${scope}`;
    } else if (provider.type === "Lastfm") {
      return `${endpoint}?api_key=${provider.clientId}&cb=${redirectUri}`;
    } else if (provider.type === "Shopify") {
      return `${endpoint}?client_id=${provider.clientId}&redirect_uri=${redirectUri}&scope=${scope}&state=${state}&grant_options[]=per-user`;
    } else if (provider.type === "Twitter" || provider.type === "Fitbit") {
      return `${endpoint}?client_id=${provider.clientId}&redirect_uri=${redirectUri}&state=${state}&response_type=code&scope=${scope}&code_challenge=${codeChallenge}&code_challenge_method=S256`;
    } else if (provider.type === "MetaMask") {
      return `${redirectUri}?state=${state}`;
    } else if (provider.type === "Web3Onboard") {
      return `${redirectUri}?state=${state}`;
    }
  }

  const app = Vue.createApp({
    data() {
      return {
        accountList: [
        {
                img: "static/img/social_github.png",
                platform: "Github",
                desc: "website: https://github.com",
                bind: false,
                provider: authInfo.GitHub,
                url: getAuthUrl(providers[0], "link")
            },
            {
                img: "static/img/social_twitter.png",
                platform: "Twitter / X",
                desc: "website: https://twitter.com",
                bind: false,
                provider: authInfo.Twitter,
                url: getAuthUrl(providers[1], "link")
            },
            {
                img: "static/img/social_facebook.png",
                platform: "Facebook",
                desc: "website: https://facebook.com",
                bind: false,
                provider: authInfo.Facebook,
                url: getAuthUrl(providers[2], "link")
            },
            {
                img: "static/img/social_wechat.png",
                platform: "Wechat",
                desc: "website: https://weixin.qq.com",
                bind: false,
                provider: authInfo.WeChat,
                url: getAuthUrl(providers[3], "link")
            }
        ]
      };
    },
    methods: {
        
    },
  });
  app.use(naive);
  app.config.compilerOptions.delimiters = ["${", "}"];
  app.mount("#app");
