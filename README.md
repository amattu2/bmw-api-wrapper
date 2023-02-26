# Introduction

This is a exploration project into the bmwusa.com ownership site internal APIs.

<details>
   <summary>/forward</summary>

   <https://mygarage.bmwusa.com/bin/api/forward?groupId=LDwJ4ZVsyNQGxmwoQ%2BJRpA%3D%3D&targetURL=%2Fmybmw%2Fresources%2Fvehicleprofile%2FgetvehicleFeatures&vin={VIN}&colorCode=0475&upholsteryCode=LCD1&lineMakeId=1>

   Decode BMW build options. This is a protected endpoint.

   Example response

   ```JSON
   {
      "responseContent":{
         "responseCode":"200",
         "responseMessage":"SUCCESS"
      },
      "errorContent":{
         "anyErrors":false,
         "errors":[

         ]
      },
      "dataContent":{
         "VehicleDetails":[
            {
               "exteriorColorCode":"0475",
               "interiorCode":"LCD1",
               "exteriorColorDesc":"Black Sapphire Metallic",
               "exteriorColorImage":"https://cache.bmwusa.com/cosy.arox?client=byo&brand=wbbm&view=paint_200x75&paint=P0475",
               "interiorColorDesc":"Coral Red/Black Dakota Leather",
               "interiorColorImage":"https://cache.bmwusa.com/cosy.arox?client=byo&brand=wbbm&view=fabric_200x75&fabric=FLCD1",
               "packageDetails":[
                  {
                     "packageName":"M Sport Package",
                     "price":3250,
                     "doNotExpand":false,
                     "options":[
                        "18\" Light alloy Star-spoke wheels style 193M-with performance run-flat tires",
                        "Aerodynamic kit",
                        "Anthracite headliner",
                        "Black Sapphire Metallic",
                        "Glacier Silver Aluminum trim",
                        "Increased top speed limiter",
                        "M steering wheel",
                        "Shadowline exterior trim",
                        "Sport seats"
                     ]
                  },
                  {
                     "packageName":"M Sport Package",
                     "price":3250,
                     "doNotExpand":false,
                     "options":[
                        "Park Distance Control (rear)"
                     ]
                  },
                  {
                     "packageName":"Premium Package",
                     "price":2650,
                     "doNotExpand":false,
                     "options":[
                        "Auto-dimming interior and exterior mirrors",
                        "Auto-dimming rearview mirror",
                        "Interior mirror with compass",
                        "Lumbar support",
                        "Universal garage-door opener"
                     ]
                  },
                  {
                     "packageName":"Added options",
                     "price":-1,
                     "doNotExpand":false,
                     "options":[
                        "BMW Assist with Bluetooth",
                        "Black Sapphire Metallic",
                        "M Sport Package",
                        "Park Distance Control (rear)",
                        "Premium Package"
                     ]
                  },
                  {
                     "packageName":"Standard Features",
                     "price":-2,
                     "doNotExpand":false,
                     "options":[
                        "BMW TeleServices",
                        "CO2 control",
                        "DVD Area coding (North America)",
                        "Dual Zone Auto Climate Control",
                        "Dynamic Cruise Control",
                        "Floor mats",
                        "Lights Package",
                        "Moonroof",
                        "Online Information Services",
                        "Real Time Traffic Information",
                        "Satellite radio preparation",
                        "Stronger electricity supply",
                        "TeleService Activation",
                        "TeleService Control",
                        "Tire pressure monitor",
                        "Voice-command",
                        "Warning triangle"
                     ]
                  }
               ]
            }
         ]
      }
   }
   ```

</details>

<details>
   <summary>/warranty-and-service</summary>

   <https://mygarage.bmwusa.com/bin/mybmw/warranty-and-service?groupId=GROUP_ID&vin={VIN}>

   Returns information relating to the vehicle's warranty and service information.

   ```JSON
   {
      "vinResponse":{
         "tcuInfo":{
            "tcuConveniencePlanExpireDate":{
               "link":false,
               "value":"00/00/0000"
            },
            "tcuExpireDate":{
               "link":false,
               "value":"00/00/0000"
            },
            "tcuMobileDirectoryNumber":"",
            "tcuActivationStatus":"",
            "tcuMobileIdentNumber":"",
            "tcuSerialNumber":""
         },
         "contract":[

         ],
         "eligibleContract":[
            {
               "expireOdometer":{
                  "expireOdometer":"999,999",
                  "odometerUnits":"M"
               },
               "contractType":"MP",
               "vendor":"MP",
               "msrp":"199.00",
               "deductible":"",
               "expireAge":"36",
               "program":{
                  "programDescription":"UC OIL 36 MO/UNL MLS",
                  "code":"0000000500"
               }
            }
         ],
         "warranty":{
            "warrantyTermDate":"MM/DD/YYYY",
            "warrantyTermAge":"48",
            "warrantyTermOdoMeter":{
               "warrantyTermOdoMeter":"50,000",
               "odometerUnits":"M"
            },
            "mfgWarrantyCoverage":[
               {
                  "contractStatus":"EXPIRED",
                  "expireOdometer":{
                     "expireOdometer":"50,000",
                     "odometerUnits":"M"
                  },
                  "contractType":"WA",
                  "expireAge":"48",
                  "expireDate":"MM/DD/YYYY",
                  "program":{
                     "programDescription":"New Vehicle Limited Warranty"
                  },
                  "enrollDate":"",
                  "effectiveDate":"MM/DD/YYYY"
               },
               {
                  "contractStatus":"EXPIRED",
                  "expireOdometer":{
                     "expireOdometer":"80,000",
                     "odometerUnits":"M"
                  },
                  "contractType":"FE",
                  "expireAge":"96",
                  "expireDate":"MM/DD/YYYY",
                  "program":{
                     "programDescription":"Federal Emissions 8/80 FE 96 MO/80K MLS"
                  },
                  "enrollDate":"MM/DD/YYYY",
                  "effectiveDate":"MM/DD/YYYY"
               },
               {
                  "contractType":"CA",
                  "program":{
                     "dcsnetUrl":"https://dealerserviceclaimsus-gf4.bmwgroup.net/WKPEmissionPortal/ng/#emissionportal",
                     "programDescription":"California Emissions",
                     "message":"Vehicle may qualify for California Emissions Coverage.  Please verify if the vehicle was sold or is registered in a California Emissions Covered State (AZ,CA,CT,DE,MA,MD,ME,NJ,NM,NY,OR,PA,RI,VT,WA). Please confirm model year, selling state and current registration of the vehicle  and  click below link to determine coverage.",
                     "infonetUrl":"https://dealerserviceclaimsus-gf4.bmwgroup.net/WKPEmissionPortal/ng/#emissionportal"
                  }
               }
            ]
         },
         "retailInfo":{
            "centerId":"",
            "currentAge":"",
            "address":{
               "city":"",
               "phone":"",
               "address1":"",
               "postalCode":"",
               "state":""
            },
            "locationId":"01",
            "retailDate":"MM/DD/YYYY",
            "inServiceDate":"MM/DD/YYYY",
            "centerName":""
         },
         "vehicle":{
            "vinMsrp":"50620.0",
            "naModelCodeDesc":{
               "code":"0933",
               "value":"335i Coupe"
            },
            "interiorColor":{
               "code":"LCD1",
               "value":"Coral Red/Black Dakota Leather"
            },
            "exteriorColor":{
               "code":"475",
               "value":"Black Sapphire Metallic"
            },
            "agModelCodeDesc":{
               "code":"WB73",
               "value":"335i Coupe"
            },
            "vehicleProdDate":"YYYY/MM",
            "lineMakeDesc":{
               "code":"01",
               "value":"BMW Car"
            },
            "vin":"",
            "chassis":"",
            "modelYear":"2009",
            "motorType":"N54 ",
            "engineSeries":"E92 "
         },
         "wholeSaleInfo":{
            "centerId":"",
            "address":{
               "city":"",
               "phone":"",
               "address1":"",
               "postalCode":"",
               "state":" "
            },
            "locationId":"",
            "wholeSaleDate":"MM/DD/YYYY",
            "centerName":""
         },
         "status":null
      },
      "overallStatus":{
         "statusDescription":"COMPLETED SUCCESSFULLY",
         "sclass":"I",
         "statusCode":"0"
      }
   }
   ```

</details>

# Credits

- <https://github.com/bimmerconnected/bimmer_connected>
