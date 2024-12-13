# Thinkport API

Get data from Thinkport GmbH, Germany. It includes the following services:

| Service                 | Type   | URL                                                                   |
|-------------------------|--------|-----------------------------------------------------------------------|
| Members                 | array  | <https://production-helloworld-5w4i.encr.app/members>                 |
| Members count           | number | <https://production-helloworld-5w4i.encr.app/members/count>           |
| Members count engineers | number | <https://production-helloworld-5w4i.encr.app/members/count/engineers> |
| Member                  | object | <https://production-helloworld-5w4i.encr.app/member/Alice>            |
| Product - Trainings     | array  | <https://production-helloworld-5w4i.encr.app/product/trainings>       |
| Locations               | array  | <https://production-helloworld-5w4i.encr.app/about/locations>         |

More request examples can be found in the 
[docs/requests/Production.http](docs/requests/Production.http).

## Usage examples

### Members

#### Example with JavaScript

```javascript
const axios = require('axios');

axios.get('https://production-helloworld-5w4i.encr.app/members')
    .then((response) => {
        console.log(response.data);
    });
```

#### Example with CURL

```bash
curl https://production-helloworld-5w4i.encr.app/members
curl https://production-helloworld-5w4i.encr.app/product/trainings
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Deployment

```bash
git push origin main
```

## Development

```bash
encore run
```

### Testing

```bash
encore test ./...
```

## License

[MIT](LICENSE)
