package main
import "flag"
import "fmt"
import "os"
import "io"
import "strconv"
import "bufio"
import "time"
import "sorter/src/algorithms/bubblesort"
import "sorter/src/algorithms/qsort"


var infile *string =flag.String("i","infile","File contains value for sorting")
var outfile *string =flag.String("o","outfile","File to recevie sorted values")
var algorithm *string =flag.String("a","qsort","sort algorithm")

func readValues(infile string)(values[]int,err error){
	file,err:=os.Open(infile)
	if err!=nil{
		fmt.Println("Faild to open the input file",infile)
		return
	}
	defer file.Close();
	br := bufio.NewReader(file)
	values = make([]int ,0)
	for{
		line,isPrefix,err1:=br.ReadLine();
		if err1!=nil{
			if err1!=io.EOF{
				err=err1
			}
		break
		}
		if isPrefix{
			fmt.Println("A too long line,seem unexpected")
			return
		}
		str:=string(line)
		value,err1:=strconv.Atoi(str)
		if err1 !=nil{
			err=err1
			return
		}
		values=append(values,value)
	}
	return
}
func writeValues(values []int,outfile string)error{
	file,err:=os.Create(outfile)
	if err!=nil{
		fmt.Println("Failed to create the output file",outfile)
		return err
	}
	defer file.Close()
	for _,value:=range values{
		str:=strconv.Itoa(value)
		file.WriteString(str+"\n")
	}
	return nil
}

func main(){
	flag.Parse()
	if infile!=nil{
		fmt.Println("infile=",*infile,"outfile=",*outfile,"algorithm=",*algorithm)
	}
	values,err:=readValues(*infile)
	fmt.Println("Read values:",values)
	if err==nil{
		t1:=time.Now()
		switch *algorithm{
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm",*algorithm,"is either unkknown or unsupported")
		}
		t2:=time.Now()
		fmt.Println("The sorting process costs",t2.Sub(t1),"to complete.")
		writeValues(values,*outfile)
		fmt.Println("writer values:",values)
	}else{
		fmt.Println(err)
	}
}