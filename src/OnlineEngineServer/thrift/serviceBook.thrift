namespace java service.demo
namespace go service.demo

struct WorkBook{
    1:required string name
    2:required i32 sheetCount
}

service WorkBooks{
    WorkBook open(1:string path)
}