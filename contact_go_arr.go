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

    status int; // 是否存储了
} ;

type UserManager struct {
	userInfoArr  []UserInfo; // 保存,用户信息的数组
	
	index int;// 下一个,可能存储的下标
	num int; // 当前的人数 
	maxSize int; // 最大的人数
};

// 初始化
func (userManager *UserManager)init_manager(maxSize int) {

    userManager.num = 0;// 当前的数量
	userManager.maxSize = maxSize;

	// 创建,最大的空间
	userManager.userInfoArr = make([]UserInfo , maxSize);
}


// 添加用户
func (userManager *UserManager) add_user() {

    // 获取,存储的下标
    index := userManager.getSaveIndex();
    if index == -1 {
        fmt.Printf("已经满了\n");
        return;
    }

    // 获取,插入的节点值
	saveNode := &userManager.userInfoArr[index];

	// 新建,终端读取器
	reader := bufio.NewReader(os.Stdin);
	
	fmt.Print("请输入名字: ");
	saveNode.name ,_ = reader.ReadString('\n')
    saveNode.name = strings.TrimRight(saveNode.name, "\n");

	fmt.Print("请输入年龄: ");
	ageStr,_ := reader.ReadString('\n')

    // fmt.Printf("%d\n", ageStr);
    ageStr = strings.TrimRight(ageStr, "\n");
    saveNode.age,_ = strconv.Atoi(ageStr)
    // saveNode.age = 44444;

	fmt.Print("请输入手机号: ");
	saveNode.phone ,_ = reader.ReadString('\n')
    saveNode.phone = strings.TrimRight(saveNode.phone, "\n");
    

    saveNode.status = 1; // 已插入
    
    // 记录下: 下一个可能,被存储的位置
    userManager.index = (userManager.index + 1) % userManager.maxSize;

    userManager.num++;

    // fmt.Println("保存的信息：", saveNode)

    fmt.Printf("添加成功\n");
    reader.ReadRune(); // 读取一个字符
}

func (userManager *UserManager)is_full_arr() bool {
    if userManager.num == userManager.maxSize {
        return true;
    } else {
        return false;
    }

}

// 获取,可以存储的下标位置
func (userManager *UserManager)getSaveIndex() int {
    if userManager.is_full_arr() {
        return -1;
    }

    index := userManager.index;

    for i := 0; i< userManager.maxSize; i++ {
        if (userManager.userInfoArr[index].status == 0) {
            return index; // 该位置, 没有被存储
        }
        index = (userManager.index + 1) % userManager.maxSize;
    }
    return -1;
}

func (userManager *UserManager) list_user() {
    if (userManager.num == 0) {
        fmt.Printf("没有用户\n");
    
        bufio.NewReader(os.Stdin).ReadRune();
        return
    }

    j := 0;
    userInfoArr := userManager.userInfoArr;

    for i:=0; i<userManager.maxSize ; i++ {
        if userInfoArr[i].status == 1 {
            fmt.Printf("姓名: %s  年龄: %d 手机号: %s\n", userInfoArr[i].name,  userInfoArr[i].age,  userInfoArr[i].phone )
            j++;
        }
        if (j == userManager.num ) {// 总共num个,已经找出来了
            break;
        }
    } // end of for

    bufio.NewReader(os.Stdin).ReadRune();

}

func (userManager *UserManager) delete_user() {

    fmt.Printf("请输入,要删除的名字: ");
    
    reader := bufio.NewReader(os.Stdin);
    name,_ := reader.ReadString('\n');

    name = strings.TrimRight(name, "\n"); // 去除 空格

    userInfoArr := userManager.userInfoArr;

    // 找到
    findFlag := 0;

    for i:=0; i < userManager.maxSize; i++ {
        if (name == userInfoArr[i].name) {
            userInfoArr[i].status = 0; // 删除
            findFlag = 1;

            userManager.num--
            break;
        }
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

    userManager.init_manager(10);
    userManager.PageUi();
    

}

