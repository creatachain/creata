module.exports = {
  theme: "creatachain",
  title: "Creatachain Hub",
  base: process.env.VUEPRESS_BASE || "/",
  themeConfig: {
    docsRepo: "creatachain/creata",
    docsDir: "docs",
    editLinks: true,
    label: "hub",
    topbar: {
      banner: true,
    },
    sidebar: {
      nav: [
        {
          title: "Resources",
          children: [
            {
              title: "Tutorials",
              path: "https://tutorials.creata.network"
            },
            {
              title: "SDK API Reference",
              path: "https://godoc.org/github.com/creatachain/creata-sdk"
            },
            {
              title: "REST API Spec",
              path: "https://creata.network/rpc/"
            }
          ]
        }
      ]
    },
    gutter: {
      editLink: true,
    },
    footer: {
      question: {
        text: "Chat with Creata developers in <a href='https://discord.gg/vcExX9T' target='_blank'>Discord</a> or reach out on the <a href='https://forum.creata.network/c/creata-sdk' target='_blank'>SDK Developer Forum</a> to learn more."
      },
      logo: "/logo-bw.svg",
      textLink: {
        text: "creata.network",
        url: "https://creata.network"
      },
      services: [
        {
          service: "medium",
          url: "https://blog.creata.network/"
        },
        {
          service: "twitter",
          url: "https://twitter.com/creata"
        },
        {
          service: "linkedin",
          url: "https://www.linkedin.com/company/creatachain/"
        },
        {
          service: "github",
          url: "https://github.com/creatachain/gaia"
        },
        {
          service: "reddit",
          url: "https://reddit.com/r/creatanetwork"
        },
        {
          service: "telegram",
          url: "https://t.me/creataproject"
        },
        {
          service: "youtube",
          url: "https://www.youtube.com/c/creataProject"
        }
      ],
      smallprint:
        "This website is maintained by Augusteum Inc. The contents and opinions of this website are those of Augusteum Inc.",
      links: [
        {
          title: "Documentation",
          children: [
            {
              title: "Creata SDK",
              url: "https://docs.creata.network"
            },
            {
              title: "Creatachain Hub",
              url: "https://hub.creata.network/"
            },
            {
              title: "Augusteum Core",
              url: "https://docs.augusteum.com/"
            }
          ]
        },
        {
          title: "Community",
          children: [
            {
              title: "Creata blog",
              url: "https://blog.creata.network/"
            },
            {
              title: "Forum",
              url: "https://forum.creata.network/"
            }
          ]
        },
        {
          title: "Contributing",
          children: [
            {
              title: "Contributing to the docs",
              url:
                "https://github.com/creatachain/gaia/blob/main/docs/DOCS_README.md"
            },
            {
              title: "Source code on GitHub",
              url: "https://github.com/creatachain/gaia/"
            }
          ]
        }
      ]
    }
  },
  plugins: [
    [
      "@vuepress/google-analytics",
      {
        ga: "UA-51029217-2"
      }
    ],
    [
      "sitemap",
      {
        hostname: "https://hub.creata.network"
      }
    ]
  ]
};
