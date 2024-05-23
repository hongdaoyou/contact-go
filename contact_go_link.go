package main

// import "./a111/a1.go"
import (
	// "./a1"
	"fmt"
	"os"
    "os/exec"
	"bufio"
    "strings"
    "strconv"
)


// 用户的信息
type UserInfo struct {
	name string;
	age int;
	phone string;

	next *UserInfo ;
} ;

type UserManager struct {
	userHead UserInfo; // 保存,用户的头
	
	num int; // 当前的人数 
};

// 初始化
func (userManager *UserManager)init_manager() {

    userManager.num = 0;// 当前的数量

}


// 添加用户
func (userManager *UserManager) add_user() {
    
    newNode := new(UserInfo);

    // 新建,终端读取器
	reader := bufio.NewReader(os.Stdin);
	
	fmt.Print("请输入名字: ");
	newNode.name ,_ = reader.ReadString('\n')
    newNode.name = strings.TrimRight(newNode.name, "\n");

	fmt.Print("请输入年龄: ");
	ageStr,_ := reader.ReadString('\n')

    // fmt.Printf("%d\n", ageStr);
    ageStr = strings.TrimRight(ageStr, "\n");
    newNode.age,_ = strconv.Atoi(ageStr)
    // newNode.age = 44444;

	fmt.Print("请输入手机号: ");
	newNode.phone ,_ = reader.ReadString('\n')
    newNode.phone = strings.TrimRight(newNode.phone, "\n");
    
    
    userHead := &userManager.userHead;

    for userHead.next != nil {
        userHead = userHead.next;
    }

    // 插入,新节点
    userHead.next = newNode;

    userManager.num++;
    
    fmt.Printf("添加成功\n");
    reader.ReadRune(); // 读取一个字符
}

func (userManager *UserManager) list_user() {
    if (userManager.num == 0) {
        fmt.Printf("没有用户\n");
    
        bufio.NewReader(os.Stdin).ReadRune();
        return
    }

    // 找出,第一个节点
    userHead := userManager.userHead.next;

    for userHead != nil {
        fmt.Printf("姓名: %s  年龄: %d 手机号: %s\n", userHead.name, userHead.age, userHead.phone )
        userHead = userHead.next
    }

    bufio.NewReader(os.Stdin).ReadRune();

}

func (userManager *UserManager) delete_user() {

    fmt.Printf("请输入,要删除的名字: ");
    
    reader := bufio.NewReader(os.Stdin);
    name,_ := reader.ReadString('\n');

    name = strings.TrimRight(name, "\n"); // 去除 空格

    // 找到
    findFlag := 0;


    // 找出,第一个节点
    userHead := &userManager.userHead;

    for userHead.next != nil {
        if userHead.next.name == name {
            userHead.next = userHead.next.next
            findFlag = 1 // 找到
            break
        }
        userHead = userHead.next
    }

    if (findFlag == 0) {
        fmt.Printf("no find\n")
    } else {
        fmt.Printf("删除成功\n")
    }

    reader.ReadRune();
}


func (userManager *UserManager)PageUi() {
    // reader := bufio.NewReader(Os.Stdin)

    // 遍历
    for {
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()

        fmt.Println("        通讯录")
        fmt.Println("    1) 添加用户")
        fmt.Println("    2) 列出用户")
        fmt.Println("    3) 删除用户")
        fmt.Println("    4) 退出\n")

        // fmt.Printf("最多的人数 %d\n", userManager.maxSize );

        inputNum := 0; // 输入的值
        exitFlag := 0

        fmt.Printf("请输入编号: ")

        fmt.Scan(&inputNum )

        // fmt.Printf("input: %d\n", inputNum);
        
        switch inputNum {
        case 1:
            userManager.add_user()
        
        case 2:
            userManager.list_user()

        case 3:
            userManager.delete_user()
        
        case 4:
            exitFlag = 1

        default:
            fmt.Println("输入错误");

            // 读取 getchar()
            reader := bufio.NewReader(os.Stdin)
            reader.ReadRune();
        }
        if exitFlag == 1 { // 退出
            break
        }
    } // end of
}


func main() {
	
    userManager := new(UserManager);

    userManager.init_manager();
    userManager.PageUi();
    

}

