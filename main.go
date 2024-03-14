package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"remotetest/test"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9005, "The Server Port")
)

type server struct {
	test.UnimplementedTestServer
}

func (s *server) GetTestResult(ctx context.Context, in *test.AmendableTest) (*test.AmendableResult, error) {
	if in.TestNumber == 1 {
		return &test.AmendableResult{Message: fmt.Sprintf("%s%s", "修正区块数分别为0,20,40,60,80,100时实验数据如下（单位ms）:",
			"传统安全多方计算方案:0,55.2,111.1,167.4,225.3,279.8"+
				"无权限者组随机选择用户方案:"+
				"0,47.3,92.5,137.5,185.3,231.1")}, nil
	}
	if in.TestNumber == 2 {
		return &test.AmendableResult{Message: fmt.Sprintf("%s%s", "作恶节点分别为1，10，20，40时成功率如下:",
			"传统修正方案:89.3,81.2,63.2,35.5"+
				"无权限者组随机选择用户方案:"+
				"90.4,83.5,66.1,37.6") +
			"本方案：" +
			"99.5,99.5,99.5,98.5"}, nil
	}
	if in.TestNumber == 3 {
		return &test.AmendableResult{Message: fmt.Sprintf("%s%s", "电子合同数量为0，50，100，150，200，250，300时成功率如下:",
			"方案35:0,27.7,51.2,69.9,81.4"+
				"方案36:0,39.5,63.4,86.61,95.1"+
				"本文方案：0,91.7,95.6,97.3,99.3")}, nil
	}
	if in.TestNumber == 4 {
		return &test.AmendableResult{Message: fmt.Sprintf("%s%s", "权限者组与门限为（5，3），（10，5），（15，8），（30，15）时实验数据如下（单位ms）:",
			"493.2,580.3,764,891.9")}, nil
	}

	return nil, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	test.RegisterTestServer(s, &server{})
	log.Printf("server listening at %v", listen.Addr())
	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
