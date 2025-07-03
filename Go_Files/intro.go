/*In Go, working with files involves reading from and wrirting to files on local filesystem.
 */

package main

import (
	//"bufio"
	"fmt"
	"os"
)

func main(){
	//Reading the file, so os.Open return the func, and error we need to handle that
	/* f, err := os.Open("example.txt");
	if err != nil {
		//log the error, or panic that program cant do anything
		panic(err);
	}
	fileInfo, err := f.Stat() //it return the describing the file and error if there in any in path
	if err != nil {
		panic(err);
	}
	fmt.Println("Name of file: ", fileInfo.Name());
	fmt.Println("Is it folder: ", fileInfo.IsDir());
	fmt.Println("File size in bit: ", fileInfo.Size());
	fmt.Println("File modified at: ", fileInfo.ModTime());
	*/

	//read file
	/* f, err := os.Open("example.txt");
	if err != nil {
		panic(err);
	}
	defer f.Close(); //closing the os

	
	fileInfo, err := f.Stat();
	if err != nil {
		panic(err);
	}
	//now read the file in buffer, buffer is temp location, it is array of bits
	buf := make([]byte, fileInfo.Size());
	d, err := f.Read(buf);
	if err != nil {
		panic(err);
	}

	//now i want to read the data
	for i := 0; i<len(buf); i++{
		println("data", d, string(buf[i]));
	}
	println("data", d, buf); //d is giving the size, buf is giving the array of bytes

	//Simple method to read the file
	// data, er := os.ReadFile("example.txt");
	// if er != nil {
	// 	panic(er);
	// }
	// fmt.Println(string(data));

	*/

	//--> Read Folders <--
	//dir, err := os.Open("."); //this means in this folder

	//now i want to go back one folder
	/* dir, err := os.Open("../"); //one folder back, this will return all the folder
	if err != nil {
		panic(err);
	}
	defer dir.Close();
	*/

	//fileInf, err := dir.ReadDir(2); //in this i need to return the number of file, so it will give those number of file

	//now if i will give -1 it is n <= 0. so it will return the all files
	/* fileInf, err := dir.ReadDir(-1);
	if err != nil {
		panic(err);
	}
	//fileInf is returning the slice
	for _, f1 := range fileInf{
		fmt.Println(f1.Name());
	} */


	// --->> Create File <<---
	f, err := os.Create("NewOne.txt");
	if err != nil {
		panic(err);
	}
	defer f.Close();

	//f.WriteString("Hey Go, Going Good");
	//f.WriteString("Nice Language"); //this work in append mode

	//creating the bytes slice
	bytes := []byte("Hello Golang");
	f.Write(bytes);

	//--> Read and write to another file (streaming fashion)
	/* sourceFile, err := os.Open("example.txt");
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close();
	
	destFile, err := os.Create("example2.txt");
	if err != nil {
		panic(err);
	}
	defer destFile.Close();

	//creating new reader, which read from bufio method
	reader := bufio.NewReader(sourceFile);
	writer := bufio.NewWriter(destFile);

	for {
		b, err := reader.ReadByte();
		if err != nil {
			if err.Error() != "EOF"{
				panic(err)
			}
			break;
		}
		e := writer.WriteByte(b); //passing the byte, this return the byte
		if e != nil {
			panic(e);
		}
	}
	//now flush the writer
	writer.Flush();
	*/

	// --->> Deleting the file <<----

	err1 := os.Remove("example2.txt");
	if err1 != nil {
		panic(err1);
	}

	fmt.Println("Progam Run Successfully");
}