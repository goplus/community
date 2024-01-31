# Go+ Community Architecture

## Architecture Overview

Go+ Community is a community for Go+ developers. It provides a platform for developers to share their knowledge and experience, and to communicate with each other.It contains the following modules:

- Account Module: User login and logout
- Article Package: Article CRUD
- Markdown Component: Markdown Editor
- Translate Package: Markdown and Video Translation
- Media Package: Media Management
- Share Component: Share to Twitter or Facebook
- Web Module: Web UI

## Architecture Design

Community is a web application based on `yap`, the interaction of each module is as follows:

![Interaction](../assets/01_module_interaction_2.png)

The main interaction process is as surrouned by user interaction.

User info is stored in the database, and the user info is used to verify the user's login status and parse the user info. And get the OAuth login authentication information from github twitter facebook and store the user information temporarily in a cookie.

The article module is the core of the goplus community, and it not only connects the user module, but also integrates the article translation and resource upload module. User can create, update, delete and query articles when they login, then the article will be stored in the database. The article operation involves the storage of HTML files, it also needs to be connected to Qiniu cloud storage.

The markdown component is used to edit the article, support markdown syntax and preview. Especially, our markdown editor supports goplus syntax highlighting, markdown translation and video translation, easy to paste pictures and videos.

The share component is used to share the article to Twitter or Facebook.

The Media package is used to manage resources to the cloud storage(QiNiu Cloud Storage). It can upload, delete and query resources.

The web module is the front-end of the community, which is used to display the community interface. It is based on yap template.


## Module list

- [Account Module](./02_account_module.md)
- [Article Package](./03_article_package.md)
- [Markdown Component](./04_markdown_component.md)
- [Translate Package](./05_translation_package.md)
- [Share Component](./06_share_component.md)
- [Media Package](./07_media_package.md)
- [Web Module](./08_web_module.md)
