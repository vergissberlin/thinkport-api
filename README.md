# Thinkport API

Get data from Thinkport GmbH, Germany. It includes the following services:

[x] Memebers
[x] Events
[x] News

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

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Deplyoment

```bash
git push origin main
```

## License

MIT
