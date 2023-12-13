package main

import (
	"encoding/json"
	"fmt"
	"medinfo/sdkInit"
	"medinfo/service"
	"medinfo/web"
	"medinfo/web/controller"
	"os"

)

const (
	cc_name = "simplecc"
	cc_version = "1.0.0"
)
func main() {
	// init orgs information
	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: os.Getenv("GOPATH") + "/src/medinfo/fixtures/channel-artifacts/Org1MSPanchors.tx",
		},

	}

	// init sdk env info
	info := sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    os.Getenv("GOPATH") + "/src/medinfo/fixtures/channel-artifacts/channel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    os.Getenv("GOPATH")+"/src/medinfo/chaincode/",
		ChaincodeVersion: cc_version,
	}

	// sdk setup
	sdk, err := sdkInit.Setup("config.yaml", &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}

	// create channel and join
	if err := sdkInit.CreateAndJoinChannel(&info); err != nil {
		fmt.Println(">> Create channel and join error:", err)
		os.Exit(-1)
	}

	// create chaincode lifecycle
	if err := sdkInit.CreateCCLifecycle(&info, 1, false, sdk); err != nil {
		fmt.Println(">> create chaincode lifecycle error: %v", err)
		os.Exit(-1)
	}

	// invoke chaincode set status
	fmt.Println(">> Set chaincode status through chaincode external service......")

	med := service.Medinfo{
		Name: "Vince Kingsley",
		Gender: "Male",
		Nation: "China",
		EntityID: "101",
		Place: "Hong Kong",
		BirthDay: "2002/01/01",
		EnrollDate: "Phone no.: 12345678",
		GraduationDate: "170cm",
		SchoolName: "100kg",
		Major: "Null",
		QuaType: "Null",
		Length: "Medications: Metformin",
		Mode: "Allergies: Penicillin, Peanuts.",
		Level: "Blood: A+",
		Graduation: "Sharing Status: true",
		CertNo: "111",
		Photo: "/static/photo/11.png",
	}

	serviceSetup, err := service.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk)
	if err!=nil{
		fmt.Println()
		os.Exit(-1)
	}
	msg, err := serviceSetup.SaveMed(med)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("The information was released successfully, the transaction number is:" + msg)
	}

	result, err := serviceSetup.FindMedInfoByEntityID("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var med service.Medinfo
		json.Unmarshal(result, &med)
		fmt.Println("Query information based on ID number successfully:")
		fmt.Println(med)
	}


	app := controller.Application{
		Setup: serviceSetup,
	}
	web.WebStart(app)
}