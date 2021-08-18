# Introduction
This is a exploration project into the bmwusa.com ownership site internal APIs.

# API Endpoints
## [/endpoints](https://mygarage.bmwusa.com/content/dam/mybmw/endpoints/Endpoints.json)
Returns a list of active bmwusa.com API endpoints.
```JSON
{
   "GET_GARAGE_VEHICLES_BY_GROUPID":"/mybmw/resources/garage/getGarageVehiclesByGroupId",
   "GET_USER_MOI":"/mybmw/resources/modelofinterest/v1/getUserMOI",
   "ADD_USER_MOI":"/mybmw/resources/modelofinterest/v1/addUserMOI",
   "NEW_USER_DETAILS":"/mybmw/resources/consumerprofile/nsclegit/newuserDetails",
   "LINK_GCID":"/mybmw/resources/consumerprofile/nsclegit/deleteOldAndLinkNewGCID",
   "VALIDATE_VIN_NSC":"/mybmw/resources/consumerprofile/nsclegit/validateVinAndCompleteLegitProcess",
   "RETRIEVE_VEHICLE_DETAILS":"/mybmw/resources/TrackMyBMW/RetrieveVehicleDetails",
   "FIND_VEHICLE_BY_VIN":"/mybmw/resources/managevehicles/findVehicleByVIN",
   "FIND_VEHICLE_BY_PRODNUM":"/mybmw/resources/managevehicles/findVehiclebyProdNum",
   "GET_MYBMW_MODELS":"/mybmw/resources/modelofinterest/v1/getmybmwmodels",
   "GET_MYBMW_MODEL_DETAILS":"/mybmw/resources/modelofinterest/v1/getmybmwmodelDetails",
   "GET_MYBMW_MOI_MODELS":"/mybmw/resources/modelofinterest/v1/getnewmodels",
   "REMOVE_USER_MOI":"/mybmw/resources/modelofinterest/v1/removeUserMOI",
   "SRO":"/mybmw/resources/vehicles/v1/SRO",
   "ADD_USER_PROD_VEHICLE":"/mybmw/resources/manageDrivenVehicles/addUserProdVehicle",
   "ADD_DRIVEN_VEHICLE":"/mybmw/resources/manageDrivenVehicles/addDrivenVehicle",
   "VEHICLE_PROFILE_INFO":"/mybmw/resources/vehicleprofile/getvehicleProfile",
   "VEHICLE_PROFILE_FEATURES_OPTIONS":"/mybmw/resources/vehicleprofile/getvehicleFeatures",
   "UPDATE_VEHICLE_NICKNAME":"/mybmw/resources/vehicleprofile/updateVehicleNickname",
   "GET_DEALER_SERVICE_CENTER":"/dealerlocator/v1/dealerlocator/getdealerdetails/",
   "GET_DEALER_SERVICE_CENTER_BY_CITY":"/dealerlocator/v1/dealerlocator/getdealerdetailsByCityState",
   "GET_DEALERS_BY_NAME":"/dealerlocator/v1/dealerlocator/getdealerdetailsByDealerName",
   "GET_PREF_SERVICE_DEALER_SRO":"/mybmw/resources/consumerprofile/v1/RetrievePrefServiceDealerSearch",
   "GET_PREF_SERVICE_DEALER_SRD":"/mybmw/resources/consumerprofile/v1/RetrieveSRDPrefServiceDealer",
   "GET_DEALER_DETAILS":"/dealerlocator/v1/dealerlocator/getdealerdetailsByDealerIdLocationId/",
   "UPDATE_SERVICE_CENTER_SRO":"/mybmw/resources/consumerprofile/v1/UpdatePrefServiceDealerSearch",
   "UPDATE_SERVICE_CENTER_SRD":"/mybmw/resources/consumerprofile/v1/UpdateSRDPrefServiceDealer",
   "GET_PREF_SALES_DEALER":"/mybmw/resources/managePreferredSalesDealer/retrieveSalesDealerDetails",
   "UPDATE_PREF_SALES_CENTER":"/mybmw/resources/managePreferredSalesDealer/updateSalesDealerDetails",
   "UB_GET_SERIES_MODELS":"/mybmw/resources/UltimateBenefits/getModelHistory",
   "MOI_TESTDRIVE":"/mybmw/resources/modelofinterest/v1/testdriveLead",
   "REMOVE_SRD_FROM_GARAGE":"/mybmw/resources/garage/removeSRDVehicle",
   "REMOVE_PRODVEHICLE_FROM_GARAGE":"/mybmw/resources/garage/removePRODVehicle",
   "GET_SERVICE_ALERTS":"/mybmw/resources/vehicleprofile/getVehicleAlerts",
   "HIDE_SERVICE_ALERT":"/mybmw/resources/vehicleprofile/hideVehicleAlertsInPG",
   "GET_DEALER_BUY_SELL_DETAILS":"/mybmw/resources/dealerprofile/v1/RetrieveDealerInfo/",
   "GET_VEHICLE_VIDEOS":"/mybmw/resources/vehicleprofile/getVehicleVideos",
   "MOI_VEHICLE_PROFILE_INFO":"/mybmw/resources/modelofinterest/v1/getMOIVehicleProfile",
   "GET_GARAGE_SAVED_BUILDS":"/mybmw/resources/garage/getGarageSavedBuild",
   "SAVE_GARAGE_BUILDS":"/mybmw/resources/garage/saveGarageSavedBuild",
   "GET_MODEL_DETAILS_BY_MODELCODE":"/mybmw/resources/modelofinterest/v1/getBmwModelsFromModelCode",
   "GET_WARRANTY_SERVICE":"/mybmw/resources/serviceNwarranty/v1/services",
   "REQUEST_QUOTE":"/mybmw/resources/modelofinterest/v1/getQuoteLead",
   "GET_XTIME_TOKEN":"/mybmw/resources/serviceToken/v1/getencryptedgcid",
   "GET_NOTIFICATIONS":"/mybmw/resources/userMessages",
   "UPDATE_NOTIFICATION_STATUS":"/mybmw/resources/userMessages/updateMessageStatus",
   "DELETE_NOTIFICATION":"/mybmw/resources/userMessages/deleteUserNotification",
   "LOAD_NOTIFICATIONS":"/mybmw/resources/loadnotifications",
   "GET_NOTIFICATIONS_STATUS":"/mybmw/resources/loadnotifications/status",
   "DELETE_NOTIFICATIONS_LOADER":"/mybmw/resources/loadnotifications/loader",
   "UPDATE_SEARCH_STATUS":"/mybmw/resources/managevehicles/updateSearchStatus",
   "GET_VALUE_SERVICES":"/mybmw/resources/ddr/v1/valueService"
}
```

## [/recalls](https://mygarage.bmwusa.com/bin/mybmw/getRecalls?targetVin=/vin/{VIN})
Returns a list of active recalls by BMW VIN number.

Example response
```JSON
{
   "vin":"{VIN}",
   "status":true,
   "year":2009,
   "make":"BMW",
   "model":"335i Coupe",
   "manufacturer_id":15,
   "recalls_available":false,
   "number_of_recalls":0,
   "refresh_date":"Apr 11, 2021",
   "recalls":[
      ...
   ]
}
```

Interesting note, opening the link directly above, without replacing `{VIN}` with a valid VIN, breaks the BMW API and leaks the API key as well as a NHTSA endpoint.

`
Invalid uri 'https://nhtsa.bmwgroup.com/safetyrecall/resources/v1/api/bmwusa/vin/{VIN}?bmwusa_api_key=CSj6v9LmEVVQ4e7TPHMTqr7F': escaped absolute path not valid
Cannot serve request to /bin/mybmw/getRecalls on this server
`

This above API key doesn't change on reload, so it must be some sort of internal API key (not dynamic/request based). 

Fetching content from the URL provided in the error message works when a valid VIN is present. The mygarage.bmwusa.com website seems to only wrap the return value, so there might not be any value in using this private endpoint.

## [/forward](https://mygarage.bmwusa.com/bin/api/forward?groupId=LDwJ4ZVsyNQGxmwoQ%2BJRpA%3D%3D&targetURL=%2Fmybmw%2Fresources%2Fvehicleprofile%2FgetvehicleFeatures&vin={VIN}&colorCode=0475&upholsteryCode=LCD1&lineMakeId=1)
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

## [/valueService](https://mygarage.bmwusa.com/content/dam/mybmw/value-service/valueService.json)
Unknown purpose.

Example response
```JSON
{
   "Status":1,
   "Message":"",
   "Result":{
      "Services":[
         {
            "Id":1,
            "Name":"Oil Change",
            "ImageUrl":"/content/images/valueservice/bmw/bmw-oil-change.jpg",
            "ImageDescription":"BMW Oil Change. BMW Service Technician performing oil service by pouring Original BMW Engine Oil into engine of a BMW vehicle in a BMW Service Center.",
            "Price":0,
            "AgCode":null,
            "DealerId":0,
            "Description":"<p>Your engine oil is vital &ndash; it keeps the moving parts of your engine lubricated and functioning at their best. Over time, all engine oil will break down and become contaminated by dirt and debris, which is why every vehicle requires regular oil changes.</p><p>Tailored perfectly to each BMW engine, BMW oil filters are made with synthetic fibers for maximum durability and effectiveness. They are capable of meeting the demands of synthetic oils and high oil temperatures, with a large filtration area (up to 3,800 sq. cm) to remove impurities such as engine sludge and deposits. Particles as fine as 0.005 millimeters and smaller are filtered out, keeping oil cleaner, reducing fuel consumption, and helping to extend the service life of your engine.</p><p>Work performed in this service:</p><ul><li>Removal of used motor oil and replacement with new premium synthetic Original BMW Engine Oil.</li><li>Removal of oil filter and installation of new BMW oil filter.</li><li>Eco-friendly elimination/recycling of oil and oil filters.</li><li>Service Interval Indicator or Condition Based Service is reset according to factory specifications.</li><li>Multi-point check, including brakes, power steering and air conditioning belts, air filter, and tires for wear and alignment.</li><li>Security check, including condition of safety belts and function of automatic-locking retractors, belt locks and belt buckles.</li></ul><p class=\"disclaimer\">Oil Service utilizes premium synthetic BMW oil. Oil viscosities may vary by model. Prices include oil, parts and labor. Topping off diesel emission fluid is not included.&nbsp; Taxes and additional costs may apply. Ask your BMW Center for further details.</p>"
         },
         {
            "Id":2,
            "Name":"Micro Filter Replacement",
            "ImageUrl":"/content/images/valueservice/bmw/bmw-microfilter.jpg",
            "ImageDescription":"BMW Oil Change with Microfilter Replacement. BMW Service Technician replacing Microfilter under the hood of a BMW vehicle in a BMW Service Center.",
            "Price":0,
            "AgCode":null,
            "DealerId":0,
            "Description":"<p>BMW interior micro filters promote the comfort of both the driver and passengers by filtering the air in the cabin. Interior filters are found in the air conditioning system and use mechanical and electromechanical processes to filter all the air before it enters the vehicle&rsquo;s interior. When driving in the city, through tunnels, or when caught in traffic jams, the system reduces the level of dust and harmful materials entering the cabin.</p><p>Work performed in this service:</p><ul><li>Removal of used micro filter and installation of new BMW micro filter.</li><li>Security check, including condition of safety belts and function of automatic-locking retractors, belt locks and belt buckles.</li><li>Coolant levels are checked.</li><li>Fluid level and antifreeze additive in windshield washer fluid reservoir are checked and topped off, if needed.</li><li>Check brake fluid level and the corresponding interval indicator (recommend replacement every two years, at the latest).</li></ul><p>This micro filter replacement service is performed by BMW Trained Technicians at your local BMW Center.&nbsp;</p><p class=\"disclaimer\">Prices include parts and labor. Taxes and additional costs may apply. Ask your BMW Center for further details.</p>"
         },
         {
            "Id":3,
            "Name":"Brake Pads / Sensor Replacement - Front",
            "ImageUrl":"/content/images/valueservice/bmw/bmw-brake-pad.jpg",
            "ImageDescription":"BMW Brake Pad Replacement, Front or Rear with Sensor. BMW Service Technician replacing front brake pad for a BMW vehicle on a lift in a BMW Service Center.",
            "Price":0,
            "AgCode":null,
            "DealerId":0,
            "Description":"<p>Your brake pads perform a crucial role: routinely applying friction to slow and stop a moving vehicle, so they need to be looked after with care. In addition to braking efficiency, your brake pads&rsquo; efforts can be felt throughout the vehicle, from fuel economy to precision handling. Regularly changing your brake pads is excellent for the extending the longevity of your BMW.</p><p>BMW brake pads assure optimized brake performance. These long-lasting pads provide outstanding braking results regardless of prevailing conditions. Even in extremely hot or cold weather, BMW brake pads maintain outstanding performance. They keep brakes from rubbing and squeaking while helping to reduce fading or weak braking power at high temperatures. When combined with BMW brake discs, naturally asbestos-free BMW brake pads offer the most precise match for optimal functionality in key control systems, such as the Antilock Braking System (ABS) and the Dynamic Stability Control (DSC).</p><p>Work performed in this service:</p><ul><li>Removal of front brake pads and installation of new front BMW brake pads.</li><li>Removal of used front brake pad sensor and installation of new front BMW brake pad sensor.</li><li>Cleaning of both sides of front brake calipers.</li><li>Removal and reinstallation of front wheels, with bolts tightened to specified torque.</li><li>Security check, including condition of safety belts and function of automatic-locking retractors, belt locks and belt buckles.</li><li>Coolant levels are checked.</li><li>Fluid level and antifreeze additive in windshield washer fluid reservoir are checked and topped off, if needed.</li><li>Check brake fluid level and the corresponding interval indicator (recommend replacement every two years, at the latest).</li></ul><p>This front brake pad replacement service is performed by BMW Trained Technicians at your local BMW Center.</p><p class=\"disclaimer\">Prices include parts and labor. Fluids as well as taxes may be additional.&nbsp; Ask your BMW Center for further details.&nbsp;</p>"
         },
         {
            "Id":4,
            "Name":"Brake Pads / Sensor Replacement - Rear",
            "ImageUrl":"/content/images/valueservice/bmw/bmw-brake-pad.jpg",
            "ImageDescription":"BMW Brake Pad Replacement, Front or Rear with Sensor. BMW Service Technician replacing front brake pad for a BMW vehicle on a lift in a BMW Service Center.",
            "Price":0,
            "AgCode":null,
            "DealerId":0,
            "Description":"<p>Your brake pads perform a crucial role: routinely applying friction to slow and stop a moving vehicle, so they need to be looked after with care. In addition to braking efficiency, your brake pads&rsquo; efforts can be felt throughout the vehicle, from fuel economy to precision handling. Regularly changing your brake pads is excellent for the extending the longevity of your BMW.</p><p>BMW brake pads assure optimized brake performance. These long-lasting pads provide outstanding braking results regardless of prevailing conditions. Even in extremely hot or cold weather, BMW brake pads maintain outstanding performance. They keep brakes from rubbing and squeaking while helping to reduce fading or weak braking power at high temperatures. When combined with BMW brake discs, naturally asbestos-free BMW brake pads offer the most precise match for optimal functionality in key control systems, such as the Antilock Braking System (ABS) and the Dynamic Stability Control (DSC).</p><p>Work performed in this service:</p><ul><li>Removal of rear brake pads and installation of new rear BMW brake pads.</li><li>Removal of used rear brake pad sensor and installation of new rear BMW brake pad sensor.</li><li>Cleaning of both sides of rear brake calipers.</li><li>Removal and reinstallation of rear wheels, with bolts tightened to specified torque.</li><li>Security check, including condition of safety belts and function of automatic-locking retractors, belt locks and belt buckles.</li><li>Coolant levels are checked.</li><li>Fluid level and antifreeze additive in windshield washer fluid reservoir are checked and topped off, if needed.</li><li>Check brake fluid level and the corresponding interval indicator (recommend replacement every two years, at the latest).</li></ul><p>This rear brake pad replacement service is performed by BMW Trained Technicians at your local BMW Center.</p><p class=\"disclaimer\">Prices include parts and labor. Fluids as well as taxes may be additional.&nbsp; Ask your BMW Center for further details.&nbsp;</p>"
         },
         {
            "Id":5,
            "Name":"Brake Disc / Pads / Sensor Replacement - Front",
            "ImageUrl":"/content/images/valueservice/bmw/bmw-brake-disc.jpg",
            "ImageDescription":"BMW Brake Disc Replacement, Front and / or Rear with Brake Pads and Sensor. BMW Service Technician replacing brake disc for a BMW vehicle on a lift in a BMW Service Center.",
            "Price":0,
            "AgCode":null,
            "DealerId":0,
            "Description":"<p>Your brake pads perform a crucial role: routinely applying friction to slow and stop a moving vehicle, so they need to be looked after with care. In addition to braking efficiency, your brake pads&rsquo; efforts can be felt throughout the vehicle, from fuel economy to precision handling. Regularly changing your brake pads is excellent for the extending the longevity of your BMW.</p><p>In order to assist with vehicle and occupant safety, brake discs must bear massive loads. BMW puts decades of rigorous testing and motorsports experience to work to achieve uncompromising levels of quality. Long-lasting BMW discs are calibrated exactly to match the engine, chassis, body, and aerodynamics of your specific BMW. Asbestos-free materials and a larger friction surface optimize brake performance, even at extreme temperatures. Fading and deformation in brakes are reduced through precise disc composition and production. The perfect interaction of brake disc and brake pads provide excellent brake efficiency.</p><p>Work performed in this service:</p><ul><li>Removal of used front brake pads and replacement with new front BMW brake pads.</li><li>Removal of used front brake pad sensor and installation of new front BMW brake pad sensor.</li><li>Cleaning of both sides of front brake calipers.</li><li>Removal of used front brake discs on wheel hub and installation of new front BMW brake discs.</li><li>Removal and reinstallation of front wheels, with bolts tightened to specified torque.</li><li>Security check, including condition of safety belts and function of automatic-locking retractors, belt locks and belt buckles.</li><li>Coolant levels are checked.</li><li>Fluid level and antifreeze additive in windshield washer fluid reservoir are checked and topped off, if needed.</li><li>Check brake fluid level and the corresponding interval indicator (recommend replacement every two years, at the latest).</li></ul><p>This front brake disc replacement service is performed by BMW Trained Technicians at your local BMW Center.</p><p class=\"disclaimer\">Prices include parts and labor. Fluids as well as taxes may be additional. Ask your BMW Center for further details.</p>"
         },
         {
            "Id":6,
            "Name":"Brake Disc / Pads / Sensor Replacement - Rear",
            "ImageUrl":"/content/images/valueservice/bmw/bmw-brake-disc.jpg",
            "ImageDescription":"BMW Brake Disc Replacement, Front and / or Rear with Brake Pads and Sensor. BMW Service Technician replacing brake disc for a BMW vehicle on a lift in a BMW Service Center.",
            "Price":0,
            "AgCode":null,
            "DealerId":0,
            "Description":"<p>Your brake pads perform a crucial role: routinely applying friction to slow and stop a moving vehicle, so they need to be looked after with care. In addition to braking efficiency, your brake pads&rsquo; efforts can be felt throughout the vehicle, from fuel economy to precision handling. Regularly changing your brake pads is excellent for the extending the longevity of your BMW.</p><p>In order to assist with vehicle and occupant safety, brake discs must bear massive loads. BMW puts decades of rigorous testing and motorsports experience to work to achieve uncompromising levels of quality. Long-lasting BMW discs are calibrated exactly to match the engine, chassis, body, and aerodynamics of your specific BMW. Asbestos-free materials and a larger friction surface optimize brake performance, even at extreme temperatures. Fading and deformation in brakes are reduced through precise disc composition and production. The perfect interaction of brake disc and brake pads provide excellent brake efficiency.</p><p>Work performed in this service:</p><ul><li>Removal of used rear brake pads and replacement with new rear BMW brake pads.</li><li>Removal of used rear brake pad sensor and installation of new rear BMW brake pad sensor.</li><li>Cleaning of both sides of rear brake calipers.</li><li>Removal of used rear brake discs on wheel hub and installation of new rear BMW brake discs.</li><li>Removal and reinstallation of rear wheels, with bolts tightened to specified torque.</li><li>Security check, including condition of safety belts and function of automatic-locking retractors, belt locks and belt buckles.</li><li>Coolant levels are checked.</li><li>Fluid level and antifreeze additive in windshield washer fluid reservoir are checked and topped off, if needed.</li><li>Check brake fluid level and the corresponding interval indicator (recommend replacement every two years, at the latest).</li></ul><p>This rear brake disc replacement service is performed by BMW Trained Technicians at your local BMW Center.</p><p class=\"disclaimer\">Prices include parts and labor. Fluids as well as taxes may be additional.&nbsp; Ask your BMW Center for further details.</p>"
         },
         {
            "Id":7,
            "Name":"Wiper Blade Replacement",
            "ImageUrl":"/content/images/valueservice/bmw/bmw-wiper-blade.jpg",
            "ImageDescription":"BMW Wiper Blade Replacement. BMW Service Technician replacing the driver�s side wiper blade on the front windshield of a BMW vehicle in a BMW Service Center.",
            "Price":0,
            "AgCode":null,
            "DealerId":0,
            "Description":"<p>Wiper blades are easily overlooked, until dangerous weather conditions arise and your view is obstructed. That&rsquo;s why we recommend inspecting your wiper blades frequently for wear. In addition to gradual wear from exposure to the elements, the rubber in a wiper blade can deteriorate due to road debris, chemicals, and ozone in the air. If you&rsquo;re experiencing excessive noise, streaking, or chattering, it&rsquo;s likely time for a replacement.</p><p>Work performed in this service:</p><ul><li>Replacement of wiper blades on both wiper arms.</li><li>Multi-point check, including brakes, power steering and air conditioning belts, air filter, and tires for wear and alignment.</li><li>Security check, including condition of safety belts and function of automatic-locking retractors, belt locks and belt buckles.</li><li>Coolant levels are checked.</li><li>Fluid level and antifreeze additive in windshield washer fluid reservoir are checked and topped off, if needed.</li><li>Check brake fluid level and the corresponding interval indicator (recommend replacement every two years, at the latest).</li></ul><p>This wiper blade replacement service is performed by BMW Trained Technicians at your local BMW Center.</p><p class=\"disclaimer\">Prices include parts and labor. Taxes and additional costs may apply. Ask your BMW Center for further details.</p>"
         },
         {
            "Id":8,
            "Name":"Engine Air Filter Replacement",
            "ImageUrl":"/content/images/valueservice/bmw/bmw-air-filter-replacement.jpg",
            "ImageDescription":"BMW Engine Air Filter Replacement. BMW Service Advisor standing in front of a BMW sedan in a BMW Service Center while reviewing information on a tablet computer with the vehicle�s owner.",
            "Price":0,
            "AgCode":null,
            "DealerId":0,
            "Description":"<p>Replacing an engine air filter does more than just increase airflow to the engine, it also increases fuel economy and reduces your vehicle&rsquo;s overall emissions. Engine air filters prevent abrasive materials from entering the engine&rsquo;s cylinders, where they would cause mechanical wear and contaminate your oil. In turn, a clean air filter promotes increased gas mileage and reduced emissions over time.</p><p>Work performed in this service:</p><ul><li>Removal and reinstallation of the air filter housing.</li><li>Replacement of used engine air filter and replacement with new BMW engine air filter.</li><li>Multi-point check, including brakes, power steering and air conditioning belts, air filter, and tires for wear and alignment.</li><li>Security check, including condition of safety belts and function of automatic-locking retractors, belt locks and belt buckles.</li><li>Coolant levels are checked.</li><li>Fluid level and antifreeze additive in windshield washer fluid reservoir are checked and topped off, if needed.</li><li>Check brake fluid level and the corresponding interval indicator (recommend replacement every two years, at the latest).</li></ul><p>This engine air filter replacement service is performed by BMW Trained Technicians at your local BMW Center.</p><p class=\"disclaimer\">Prices include parts and labor. Taxes and additional costs may apply. Ask your BMW Center for further details.</p>"
         },
         {
            "Id":9,
            "Name":"Spark Plug Replacement",
            "ImageUrl":"/content/images/valueservice/bmw/bmw-air-filter-replacement.jpg",
            "ImageDescription":"BMW Engine Air Filter Replacement. BMW Service Advisor standing in front of a BMW sedan in a BMW Service Center while reviewing information on a tablet computer with the vehicle’s owner.",
            "Price":0,
            "AgCode":null,
            "DealerId":0,
            "Description":"<p>Spark plugs are what start your BMW. Over time, they can wear down, making it more difficult to get the vehicle moving, especially in very cold conditions. Older spark plugs require the use of more gasoline in order to start the vehicle, which results in increased carbon emissions. Keeping your spark plugs fresh can help protect your vehicle&rsquo;s fuel economy and contribute to the overall health of your engine.</p><p>Work performed in this service:</p><ul><li>Removal and reinstallation of engine cover.</li><li>Removal and reinstallation of ignition coils.</li><li>Removal of all used spark plugs and replacement with new BMW spark plugs.</li><li>Multi-point check, including brakes, power steering and air conditioning belts, air filter, and tires for wear and alignment.</li><li>Security check, including condition of safety belts and function of automatic-locking retractors, belt locks and belt buckles.</li><li>Coolant levels are checked.</li><li>Fluid level and antifreeze additive in windshield washer fluid reservoir are checked and topped off, if needed.</li><li>Check brake fluid level and the corresponding interval indicator (recommend replacement every two years, at the latest).</li></ul><p>This spark plug replacement service is performed by BMW Trained Technicians at your local BMW Center.</p><p class=\"disclaimer\">Prices include parts and labor. Taxes and additional costs may apply. Ask your BMW Center for further details.</p>"
         },
         {
            "Id":19,
            "Name":"Coolant Pump",
            "ImageUrl":"/content/images/valueservice/bmw/bmw-oil-change.jpg",
            "ImageDescription":"",
            "Price":0,
            "AgCode":null,
            "DealerId":0,
            "Description":""
         }
      ]
   }
}
```

## [/warranty-and-service.html](https://mygarage.bmwusa.com/bin/mybmw/warranty-and-service?groupId=GROUP_ID&vin=VIN)

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
