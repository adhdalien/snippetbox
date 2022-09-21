# Snippetbox

Following the [Let's Go -book by Alex Edwards](https://lets-go.alexedwards.net).

## Routes

| Method | Pattern         | Handler       | Purpose                     | 
| ------ | --------------- | ------------- | --------------------------- |
| ANY    | /               | home          | Display the home page.      |
| ANY    | /snippet/view   | snippetView   | Display a specific snippet. |
| POST   | /snippet/create | snippetCreate | Create a new snippet.       |