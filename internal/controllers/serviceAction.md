# Synopsis
{{binaryName}} {{actionName}} \<modifier\> \<service name\>
# Description
| modifier     | description                       |
|--------------|-----------------------------------|
| list         | List available services.          |
| install      | Install service.                  |
| uninstall    | Uninstall service.                |
| health       | Show health of the service.       |
| action       | Perform an action for the service |
| list-actions | List the service actions.         |

# Examples
```
{{binaryName}} {{actionName}} install webserver
```
Install the webserver on the current target.
