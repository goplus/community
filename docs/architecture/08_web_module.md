# Web module

## Module purpose

This module is used to connect web front-end static pages to form an overall website.

## Module scope

This module uses yap to parse HTML, and renders data on the website pages.

## Module structure

![Web](../assets/08_web_module.png)

The whole module can be divided into two kind of files - HTML file and public JS file, and public JS file can be introduced through the `src` attribute of the `<script>` tag into each HTML file, so that the latter can obtain the vue.runtime.js. What's more, each HTML file can be divided into three kinds of components - header, content and footer.

### home_yap.html

**Purpose:** to present the article list, including the title, introduction and other partial information of the article.

**Interface:** none.

**Dependencies:** header_yap.html, footer_yap.html

### article_yap.html

**Purpose:** to present the article details, including the article content, the creator's infomation, comment area, and so on; to render the content of the markdown file to HTML.

**Interface:** none.

**Dependencies:** header_yap.html

### edit_yap.html

**Purpose:** to provide the editor to edit articles in markdown format; to post articles.

**Interface:** none.

**Dependencies:** header_yap.html

### user_yap.html

**Purpose:** to display the user's detailed personal information and allow users to edit; to manage (edit and delete) user's own posted articles.

**Interface**: none.

**Dependencies:** header_yap.html, footer_yap.html

### status code (5xx_yap.html, 2xx_yap.html, 4xx_yap.html)

**Purpose:** to accomplish error page redirection, supporting automatic redirect to 5xx/2xx/4xx error page.

**Interface:** none.

**Dependencies:** header_yap.html

## Module Interface

None.

## Functions

None.