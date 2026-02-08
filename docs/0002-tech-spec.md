MVP Scope:
- Single-Page Architecture: All content (About, Resume, Projects) will live on one page.
- Click to copy email: To avoid bot spam and scraping that appears on email forms. Instead of a form with a backend mail server, I plan to implement a "Click-to-Copy" component with a obfuscation technique so bots cannot read it.
- Page Data in JSON file: Will contain my project, about me, and reusme information. This will be my source of truth, that my Go application will read to populate the page.
