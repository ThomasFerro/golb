# Golb - The statically generated blogging platform

> The blogging platform that put the reader first

Golb is a static site generator for anyone wanting to get his blog up and running and focus on the essential: the content.

It aims to fully integrate with the Github ecosystem so you can host your posts and site without loosing time and motivation on technical details.

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

## Options

| Option name | Description | Required ? |
|-|-|-| 
| `posts_path` | Path to your posts | Yes |
| `blog_title` | Your blog's title | Yes |
| `locale` | Your blog's locale | Yes |
| `post_page_template_path` | Path to your post page template if you want a custom one | No |
| `home_page_template_path` | Path to your home page template if you want a custom one | No |
| `dist_path` | Path to your blog's destination (used to retrieve the generated blog and host it) | Yes |
| `global_assets_path` | Path to the assets to add to your blog | No |
