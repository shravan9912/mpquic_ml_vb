package quic
 
import (
	"fmt"
        "time"
	"os"
)


var MAaF [banditDimension][banditDimension]float64
var MAaS [banditDimension][banditDimension]float64
var MbaF [banditDimension]float64
var MbaS [banditDimension]float64
type detectionAgent struct{
	detectionbase	[]float64
}

func (det * detectionAgent) Setup(){
	det.detectionbase = make([]float64,0)
}

func (det * detectionAgent) Clear(){
	det.detectionbase = nil
}

func (det * detectionAgent) AddStep(step float64){
        det.detectionbase = append(det.detectionbase, step)
}

func (det * detectionAgent) CloseExperience(num int, path string){

		fileName := fmt.Sprintf(path + "detectionAgent_%d.csv", num)

		file, err := os.Create(fileName)
		if err != nil{
			panic(err)
		}
		defer file.Close()
                fmt.Fprintf(file, "PathInfo\n")
		for _, element := range det.detectionbase{
                        fmt.Fprintf(file, "%.4f\n", element)
		}
}


 
func main() {
        var dumpDetection detectionAgent
        dumpDetection.Setup()
        start := time.Now()
        for z := 0; z < 3; z++ {
        for i := 0; i < 600; i++ {
	    dumpDetection.AddStep(float64(i))
	}
        dumpDetection.CloseExperience(int(z),DetectAgentPath) 
        dumpDetection.Clear()
        }    
        elapsed := time.Since(start)
        fmt.Printf("Binomial took %s", elapsed) 
	//os.Remove("/Users/hongjiawu/Documents/bitbuck/mosaic-classification-based-scheduling/MAaF")
	os.Create("/home/shravan/Documents/dronenextwork/new/mpquic_ml_vb/ool/lin/lin")
	file2, _ := os.OpenFile("/home/shravan/Documents/dronenextwork/new/mpquic_ml_vb/ool/lin/lin", os.O_WRONLY, 0600)
	for i := 0; i < banditDimension; i++ {
		for j := 0; j < banditDimension; j++ {
			if i == j{
				fmt.Fprintf(file2, "%.4f\n", 1.0)
			} else {
				fmt.Fprintf(file2, "%.4f\n", 0.0)
			}
		}
	}
	for i := 0; i < banditDimension; i++ {
		for j := 0; j < banditDimension; j++ {
			if i == j{
				fmt.Fprintf(file2, "%.4f\n", 1.0)
			} else {
				fmt.Fprintf(file2, "%.4f\n", 0.0)
			}
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < banditDimension; j++ {
			if i == j{
				fmt.Fprintf(file2, "%.4f\n", 0.0)
			}
		}
	}
	file2.Close()
	file, err := os.Open("/home/shravan/Documents/dronenextwork/new/mpquic_ml_vb/ool/lin/lin")
	if err != nil {
    	panic(err)
	}

	for i := 0; i < banditDimension; i++ {
		for j := 0; j < banditDimension; j++ {
			fmt.Fscanln(file, &MAaF[i][j])
			fmt.Printf("%.4f\n",MAaF[i][j])
		}
	}

	for i := 0; i < banditDimension; i++ {
		for j := 0; j < banditDimension; j++ {
			fmt.Fscanln(file, &MAaS[i][j])
			fmt.Printf("%.4f\n",MAaS[i][j])
		}
	}

	for i := 0; i < banditDimension; i++ {
		fmt.Fscanln(file, &MbaF[i])
		fmt.Printf("%.4f\n",MbaF[i])
	}

	for i := 0; i < banditDimension; i++ {
		fmt.Fscanln(file, &MbaS[i])
		fmt.Printf("%.4f\n",MbaS[i])
	}
	file.Close()
}
