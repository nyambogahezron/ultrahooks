import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "ultrahooks docs",
  description: "ultrahooks docs",
  base: '/ultrahooks/',
  head: [['link', { rel: 'icon', href: '/ultrahooks/favicon.png' }]],
  themeConfig: {
    logo: '/logo.png',
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Documentation', link: '/README' },
      { text: 'Changelog', link: '/changelog' }
    ],

    sidebar: [
      {
        text: 'Introduction',
        items: [
          { text: 'Getting Started', link: '/README' },
          { text: 'Installation', link: '/installation' },
          { text: 'Features & Capabilities', link: '/features' },
          { text: 'CLI Commands', link: '/commands' }
        ]
      },
      {
        text: 'Deep Dives',
        items: [
          { text: 'Architecture', link: '/architecture' },
          { text: 'Git Hooks Masterclass', link: '/git-hooks-course' }
        ]
      },
      {
        text: 'Project',
        items: [
          { text: 'Contributing', link: '/contributing' },
          { text: 'Changelog', link: '/changelog' }
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/nyambogahezron/ultrahooks' }
    ]
  }
})
