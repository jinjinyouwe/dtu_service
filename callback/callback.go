package callback

import(
    "fmt"
    "log"
)

func DtuIsAvalid(dtuid string) bool{

    log.Printf("dtu %s is avalid\n", dtuid)
    return true
}

func UploadDtuMsg(dtuid string, data []byte)bool{

    if data[0] == 0xFF{
        log.Printf("%s keep alive msg\n", dtuid)
        return true
    }

    msg := fmt.Sprintf("%X",data)
    log.Printf("upload dtu: %s msg: %s\n", dtuid, msg)
    return true

}


func UploadDtuMsgHook(dtuid string, data []byte) bool{

    if len(data) > 1 && data[0] == 0xFF{
        log.Printf("%s keep alive msg\n", dtuid)
        return true
    }
    if len(data) > 6 && data[5] == 0xC0{
        UploadDtuMsg(dtuid, data)
        return true
    }

    return false
}

func UploadServerMsg(msg string) bool{

    log.Printf("upload service msg %s\n", msg)
    return true
}

func init(){
    log.Println("callback running...")
}


