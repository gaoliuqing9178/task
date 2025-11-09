package service

import (
	"errors"
	"strconv"
	"time"
	"todolist/config"
	"todolist/internal/dto"
	"todolist/internal/model"
	"todolist/internal/repository"
	"todolist/internal/vo"
	"todolist/pkg/jwtUtil"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Login(req *dto.LoginReq) (*vo.UserVO, string, error) {
	user, err := s.userRepository.GetByAccount(req.Account)
	if err != nil {
		return nil, "", errors.New("账号不存在")
	}

	if user.Password != req.Password {
		return nil, "", errors.New("密码错误")
	}

	userVO := &vo.UserVO{
		Account: user.Account,
		Name:    user.Name,
	}

	// 生成token
	token, err := jwtUtil.GenerateToken(strconv.FormatUint(uint64(user.ID), 10), time.Duration(config.Get().Jwt.Expire)*time.Second, config.Get().Jwt.SecretKey)
	if err != nil {
		return nil, "", err
	}

	return userVO, token, nil
}

func (s *UserService) Register(req *dto.RegisterReq) (*vo.UserVO, error) {
	_, err := s.userRepository.GetByAccount(req.Account)
	if err == nil {
		return nil, errors.New("用户已存在")
	}

	user := &model.User{
		Name:     req.Name,
		Account:  req.Account,
		Password: req.Password,
	}

	err = s.userRepository.Create(user)
	if err != nil {
		return nil, errors.New("注册失败")
	}

	userVO := &vo.UserVO{
		Account: user.Account,
		Name:    user.Name,
	}

	return userVO, nil
}

func (s *UserService) UpdateUserInfo(req *dto.UpdateUserInfoReq, id uint) (*vo.UserVO, error) {
	userOld, err := s.userRepository.GetByID(id)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	user := &model.User{
		Name:     req.Name,
		Account:  userOld.Account,
		Password: req.Password,
	}
	user.ID = id

	if req.Name == "" {
		user.Name = userOld.Name
	}
	if req.Password == "" {
		user.Password = userOld.Password
	}

	err = s.userRepository.Update(user)
	if err != nil {
		return nil, errors.New("更新失败")
	}

	userVO := &vo.UserVO{
		Account: user.Account,
		Name:    user.Name,
	}

	return userVO, nil
}
