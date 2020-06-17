# IBC Notes

How to integrate. Looking at `x/transfer` for example:

* `IBCModule` interface is defined in [05-port](./x/ibc/05-port/types/modules.go)
* Simple implementation of those interfaces is defined in [ibc-transfer](./x/ibc-transfer/module.go#L181)

Configuring in `app.go`:

* You must register [IBC module](./app/app.go#L102) as well as your custom module
* There are some funny [capability keeper](./app/app.go#L219-L222) to control access to IBC
* Which must be [passed to the ibc-enabled module](./app/app.go#L275-L279) along with `IBCKeeper.ChannelKeeper` and `IBCKeeper.PortKeeper`
* This [CapabilityKeeper](./app/app.go#L303) must also be registered on the app

Let's see how the capabilities work:

* ?????? utter confusion from `x/capabilities` source code

