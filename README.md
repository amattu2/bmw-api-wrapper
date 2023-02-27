# Introduction

This is a internal API exploratory project aimed at uncovering APIs exposed by BMW, Mini Cooper, & related brands.
As with any project of this type, it's a continuous work-in-progress.

If you have discovered an issue with the HTTP collection, or found an endpoint not listed or documented here, please open a PR or new Issue.

This project is focused on supporting the North American APIs, but will ideally work with european and asian APIs as well. Generally speaking, it's just a minor URL change for other markets.

*Read on to understand more about this project.*

## What is ConnectedDrive?

Almost every modern BMW, Mini, Rolls Royce, or Toyota Supra (2020+) is equipped with a product called ConnectedDrive or Connect. This featureset allows native connectivity to your vehicle through your smartphone or via other channels (websites, services, etc). In some cases, generally with newer models, ConnectedDrive allows for insights into your current vehicles status (charging/gas, mileage, location), as well as remote control (unlock/horn/remotestart).

ConnectedDrive has a plethora of hidden web-APIs that is used to receive information or dispatch commands to your vehicle remotely. The goal of this project is to explore all of those endpoints, document how to use them, and to provide a interface for doing so (i.e. Postman or Thunder Test).

Due to the fact that I can't possibly test against every single BMW on the road, it should be understood that the documentation may not be 100% accurate, but feel free to contribute any knowledge or corrections you have.

[Read more](https://www.bmwusa.com/explore/connecteddrive.html)

## What is My Garage

My Garage is a facet available across BMW brands to access information about your vehicle through a web interface. Similar to the BMW ConnectedDrive situation, each brand has it's own interface, but it's generally supported by the same backend. The My Garage service is included in this repo because it offers additional endpoints and datasets not available solely through the ConnectedDrive channels. The authentication mechanisms are identical between ConnectedDrive and My Garage.

[Mini Cooper My Garage](https://mygarage.miniusa.com/)

[BMW My Garage](https://mygarage.bmwusa.com/)

# Getting Started

## ThunderClient / Postman

There are several HTTP collections. A brief description of each is below.

- **VCS / My Garage** – The primary collection with API endpoints that relate to Vehicle Communication Service (VCS) and My Garage. They use the same authentication mechanisms.
- **Inventory** – A separate collection that holds requests relating to the used vehicle inventory searches.

### Setup

To use the Thunder Client (Visual Studio Code extension) or Postman HTTP collection, see the [.vscode](/.vscode) folder.

You can import the collection via the Thunder Client GUI.

Rename [`.env.example`](./.vscode/.env.example) to [`.env`](./.vscode/.env) and update the following placeholders to your values

- email
- password
- vin

Do not update `ocp-apim-subscription-key`. If you're using Thunder Client, link the `.env` to your Thunder Client environment.

You may also need to link the [utils.js](./.vscode/utils.js) helper functions to the Thunder Client collection. To do so,

1. Import the collection
2. Right click the "Collection" folder and press settings
3. Click the "Scripts" tab and import the `utils.js` file

### Usage

To perform the authentication process, run the Authentication folder requests in order

1. OAuth Config
2. Authenticate (1/2)
3. Authenticate (2/2)
4. Token

To perform a relogin with a refresh token (after you have already completed the above), run the following requests in order

1. OAuth Config *Optional*
2. Refresh

These requests will create or update the environment variables that are used by the rest of the collection.

## CLI Client

*To be completed*.

# Resources

The resources below were used as a reference during creation of this project. Big thanks to them.

- <https://github.com/bimmerconnected/bimmer_connected>

# Disclosure

This project is not affiliated with BMW, Mini Cooper, or any other related brand. It's likely that usage of this project is against a TOS somewhere. Use at your own risk and see the [LICENSE](./LICENSE).
