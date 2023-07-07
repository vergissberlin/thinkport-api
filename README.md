# Thinkport API

Get data from Thinkport GmbH, Germany. It includes the following services:

| Service | Type |  URL |
| --- | --- | --- |
| Members | array | <https://api.thinkport.andrelademann.de/members> |
| Member | object | <https://api.thinkport.andrelademann.de/member/Alice> |
| Product - Trainings | array | <https://api.thinkport.andrelademann.de/product/trainings> |
| Locations | array | <https://api.thinkport.andrelademann.de/about/locations> |

## Usage

### Members

JavaScript with axios:

```javascript
const axios = require('axios');

axios.get('https://api.thinkport.andrelademann.de/members')
    .then((response) => {
        console.log(response.data);
    });
```

CURL:

```bash
curl https://api.thinkport.andrelademann.de/members
curl https://api.thinkport.andrelademann.de/product/trainings
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Deplyoment

```bash
git push origin main
```

## License

[MIT](LICENSE)
