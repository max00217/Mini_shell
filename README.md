# Minishell 🐚

Go 기반의 Lua 스크립팅을 지원하는 셸 프로젝트

## 소개

- Go로 구현된 경량 쉘
- Lua 스크립팅으로 동작을 커스터마이징 가능
- ~/.mnshrc.lua 로 설정 관리  

## 지원 플랫폼

| 플랫폼 | 지원 여부 | 비고 |
|--------|----------|------|
| Linux | ✅ | 메인 개발 환경 |
| Windows | 🚧 | 추후 지원 예정 |
| macOS | 🚧 | 추후 지원 예정 |

## 기술 스택

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![Lua](https://img.shields.io/badge/Lua-2C2D72?style=flat&logo=lua&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=flat&logo=linux&logoColor=black)

## 실행 방법

```bash
git clone https://github.com/max00217/Mini_shell.git
cd Mini_shell
go run main.go
```

## 기능

- [x] 기본 REPL
- [ ] 내장 명령어 (cd, pwd, exit, history)
- [ ] Lua 엔진 연동
- [ ] Lua alias/훅
- [ ] 파이프 (`|`)
- [ ] 리다이렉션 (`>`, `<`)
- [ ] 자동완성 (Tab)
- [ ] 프롬프트 커스터마이징
- [ ] Windows 지원
- [ ] macOS 지원
- [ ] 크로스플랫폼 경로 처리

## 설정 예시 (~/.mnshrc.lua)

```lua
-- 프롬프트 설정
prompt = "mnsh λ "

-- alias 설정
alias = {
  ll = "ls -lh",
  la = "ls -lah",
  gs = "git status",
}

-- 명령어 실행 전 훅
function on_before_execute(cmd)
  -- 원하는 동작 추가
end
```

## 개발 로드맵

### 1단계 ✅
- 기본 REPL 구현
- 외부 명령어 실행

### 2단계 (진행 중)
- 내장 명령어 추가
- 히스토리 기능

### 3단계
- Lua 엔진 연동
- 설정 파일 로드

### 4단계
- 파이프/리다이렉션
- 자동완성

### 5단계 (장기 목표)
- Windows 지원
- macOS 지원
- 크로스플랫폼 경로 처리

## TODO

### 핵심 기능
- [ ] 내장 명령어 구현 (cd, pwd, exit, history)
- [ ] 환경변수 처리 ($HOME, $PATH 등)
- [ ] 히스토리 저장/불러오기

### Lua 연동
- [ ] gopher-lua 엔진 연동
- [ ] ~/.mnshrc.lua 설정 파일 로드
- [ ] Lua alias 처리
- [ ] on_before_execute 훅
- [ ] on_after_execute 훅

### 고급 기능
- [ ] 파이프 (`ls | grep foo`)
- [ ] 리다이렉션 (`>`, `<`, `>>`)
- [ ] 백그라운드 실행 (`&`)
- [ ] 자동완성 (Tab)
- [ ] 프롬프트 커스터마이징

### 크로스플랫폼
- [ ] Windows 명령어 호환성
- [ ] macOS 지원
- [ ] 크로스플랫폼 경로 처리 (filepath 패키지)
- [ ] Windows/Linux 빌드 자동화

## 참고

- [gopher-lua](https://github.com/yuin/gopher-lua)
- [Neovim](https://neovim.io) - Lua 스크립팅 참고
```