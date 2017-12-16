package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"testing"
	"time"
)

func Test(t *testing.T) {
	table := []rune{}
	for _, r := range DEFAULT_TABLE {
		table = append(table, r)
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, l := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 40, 50, 100, 200, 1000, 1001, 1002, 1003, 1004, 1005, 1006, 1007, 1008, 1009, 1010, 1011, 1012, 1013, 1014, 1015, 1016, 1017, 1018, 1019, 1020, 1021, 1022, 1023, 1024, 1025, 1026, 1027, 1028, 1029, 1030, 1031, 1032, 1033, 1034, 1035, 1036, 1037, 1038, 1039, 1040, 1041, 1042, 1043, 1044, 1045, 1046, 1047, 1048, 1049, 1050, 1051, 1052, 1053, 1054, 1055, 1056, 1057, 1058, 1059, 1060, 1061, 1062, 1063, 1064, 1065, 1066, 1067, 1068, 1069, 1070, 1071, 1072, 1073, 1074, 1075, 1076, 1077, 1078, 1079, 1080, 1081, 1082, 1083, 1084, 1085, 1086, 1087, 1088, 1089, 1090, 1091, 1092, 1093, 1094, 1095, 1096, 1097, 1098, 1099, 1100, 1101, 1102, 1103, 1104, 1105, 1106, 1107, 1108, 1109, 1110, 1111, 1112, 1113, 1114, 1115, 1116, 1117, 1118, 1119, 1120, 1121, 1122, 1123, 1124, 1125, 1126, 1127, 1128, 1129, 1130, 1131, 1132, 1133, 1134, 1135, 1136, 1137, 1138, 1139, 1140, 1141, 1142, 1143, 1144, 1145, 1146, 1147, 1148, 1149, 1150, 1151, 1152, 1153, 1154, 1155, 1156, 1157, 1158, 1159, 1160, 1161, 1162, 1163, 1164, 1165, 1166, 1167, 1168, 1169, 1170, 1171, 1172, 1173, 1174, 1175, 1176, 1177, 1178, 1179, 1180, 1181, 1182, 1183, 1184, 1185, 1186, 1187, 1188, 1189, 1190, 1191, 1192, 1193, 1194, 1195, 1196, 1197, 1198, 1199} {
		input := make([]byte, l)
		_, err := io.ReadFull(r, input)
		if err != nil {
			t.Fatal(err)
		}

		buf1 := bytes.NewBuffer(input)
		buf2 := &bytes.Buffer{}
		buf3 := &bytes.Buffer{}

		encoder := NewEncoder(table)
		decoder := NewDecoder(table)

		encoder.Encode(buf1, buf2)
		decoder.Decode(buf2, buf3)

		input1 := buf3.Bytes()
		if !compareBytesSlice(&input, &input1) {
			t.Fatalf("Fail on %d bytes:\nExpect: %s\nGot:    %s", l, bytesToString(input), bytesToString(input1))
		}
	}
}

func compareBytesSlice(s1, s2 *[]byte) bool {
	if s1 == nil || s2 == nil {
		return false
	}
	if len(*s1) != len(*s2) {
		return false
	}
	for i := 0; i < len(*s1); i++ {
		if (*s1)[i] != (*s2)[i] {
			return false
		}
	}
	return true
}

func bytesToString(b []byte) string {
	buf := bytes.Buffer{}
	for _, _b := range b {
		buf.WriteString(fmt.Sprintf("%02x ", _b))
	}
	return buf.String()
}
