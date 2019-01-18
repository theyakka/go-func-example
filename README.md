This repo contains an example of how separate "common" code from you functions code via a local module. It is written for Google Cloud Function (Go 1.11).

Please read the blog post [here](https://medium.com/yakka/cloud-functions-in-go-94c1014a6fe4) for a breakdown of how the code works.

## Deploying

To deploy the code, first vendor the local module via the following command:

```
go mod vendor
```

Then deploy the function using the `gcloud` command line tool with the command:

```
gcloud functions deploy OutputMessage --runtime go111 --trigger-http
```

You can then use the url that is printed for the `httpsTrigger` value to access the deployed cloud function.

## License

This code is released under a modified MIT license. See LICENSE for details. Use it however you want.

<hr/>
Yakka means work.
<p/>
<a href="https://theyakka.com" target="_yakka">
<img src="https://storage.googleapis.com/yakka-logos/logo_wordmark.png?1234"
  width="80"></a>