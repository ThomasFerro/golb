# Golb - The statically generated blogging platform

> The blogging platform that put the reader first

## How to use

Here is an example of how to use this action in your workflow:

```yml
name: Build blog

on:
  workflow_dispatch:

jobs:
  build-blog:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout
        uses: actions/checkout@v2
      - uses: ThomasFerro/golb@v0.1.0
        with:
          posts_path: /github/workspace/posts
          dist_path: /github/workspace/dist
          global_assets_path: /github/workspace/blog/assets
          blog_title: Thomas Ferro
          post_page_template_path: /github/workspace/blog/postPageTemplate.go.html
          home_page_template_path: /github/workspace/blog/homePageTemplate.go.html
      - name: Deploy to Github Pages
        uses: peaceiris/actions-gh-pages@v3
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: dist
```
