# go lang study

## 특징
1. constructor method가 없다.
2. class 라는 개념이 없다.
3. private public 은 대소문자로 사용한다
4. go lang 에서 Exception 은 없다
5. java , javascript 의 null 은 go lang 에서는 nil
6. 인터페이스 사용시 구현만 해주면 자동으로 호출 된다. (Package가 인터페이스에 정의된 메소드를 호출 해준다)
ex) fmt->stringer, Marshal->TextMarshaler
ex) 인터페이스 작성 예시

type TextMarshaler interface {
	MarshalText() (text []byte, err error)
}


## GC 특징 
1. 정적 유형과 동적 유형으로 나뉜다.
2. 실행파일 안에 가비지 컬렉터가 내장
3. CMS (Concurrent Mark & Sweep) 만을 수행하기에 다소 가벼운 특징 ( Go 1.10 기준) -> 현재 기준으로 조사가 필요하다
4. 비압축/비세대 원칙 고수
5. 


- 참조
https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8

## channels

## go routines
