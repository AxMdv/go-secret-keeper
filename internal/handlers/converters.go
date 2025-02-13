package handlers

import (
	"secret-keeper/internal/model"
	pb "secret-keeper/internal/proto"
)

func AllDataServiceToPb(userData model.AllUserData) *pb.GetUserStoredDataResponse {
	response := &pb.GetUserStoredDataResponse{}
	for _, logPass := range userData.LoginPassword {
		lp := LoginPassToPB(logPass)
		response.LoginPasswords = append(response.LoginPasswords, lp)
	}
	for _, textData := range userData.TextData {
		td := TextDataToPB(textData)
		response.TextData = append(response.TextData, td)
	}
	for _, binData := range userData.BinaryData {
		bd := BinDataToPB(binData)
		response.BinaryData = append(response.BinaryData, bd)
	}
	for _, bankCard := range userData.BankCard {
		bc := BankCardToPB(bankCard)
		response.BankCards = append(response.BankCards, bc)
	}
	return response
}

func LoginPassToPB(loginPass model.LoginPassword) *pb.LoginPassword {
	return &pb.LoginPassword{
		Id:       loginPass.ID,
		Login:    loginPass.Login,
		Password: loginPass.Password,
		Metadata: loginPass.Metadata,
	}
}

func TextDataToPB(textData model.TextData) *pb.TextData {
	return &pb.TextData{
		Id:       textData.ID,
		Data:     textData.Data,
		Metadata: textData.Metadata,
	}
}

func BinDataToPB(binData model.BinaryData) *pb.BinaryData {
	return &pb.BinaryData{
		Id:       binData.ID,
		Data:     binData.Data,
		Metadata: binData.Metadata,
	}
}

func BankCardToPB(bankCard model.BankCard) *pb.BankCard {
	return &pb.BankCard{
		Number:    int64(bankCard.CardNumber),
		OwnerName: bankCard.Owner,
		ExpDate:   bankCard.ExpDate,
		Cvv:       int64(bankCard.CVV),
	}
}

func PBToLoginPass(in *pb.LoginPassword) *model.LoginPassword {
	return &model.LoginPassword{
		ID:       in.Id,
		Login:    in.Login,
		Password: in.Password,
		Metadata: in.Metadata,
	}
}

func PBToTextData(in *pb.TextData) *model.TextData {
	return &model.TextData{
		ID:       in.Id,
		Data:     in.Data,
		Metadata: in.Metadata,
	}
}

func PBToBinaryData(in *pb.BinaryData) *model.BinaryData {
	return &model.BinaryData{
		ID:       in.Id,
		Data:     in.Data,
		Metadata: in.Metadata,
	}
}

func PBToBankCard(in *pb.BankCard) *model.BankCard {
	return &model.BankCard{
		CardNumber: int(in.Number),
		Owner:      in.OwnerName,
		ExpDate:    in.ExpDate,
		CVV:        int(in.Cvv),
	}
}
