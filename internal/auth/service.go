package auth

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"yasser-backend/internal/auth/jwt"
	"yasser-backend/internal/user"
	customerrors "yasser-backend/pkg/errors"
)

type Service struct {
	authRepo Repository
	userRepo user.Repository
}

func NewService(authRepo Repository, userRepo user.Repository) *Service {
	return &Service{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (s *Service) Login(phoneNumber string) error {
	otp, err := s.generateOTP()
	if err != nil {
		return err
	}
	fmt.Printf("Generated OTP for %s: %s\n", phoneNumber, otp)

	hashedOtp, err := bcrypt.GenerateFromPassword([]byte(otp), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	existingUser, err := s.userRepo.FindByPhone(phoneNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if existingUser == nil {
		_, err = s.userRepo.Create(phoneNumber)
		if err != nil {
			return err
		}
	}

	_, err = s.authRepo.SetPhoneOtp(phoneNumber, string(hashedOtp))
	if err != nil {
		return err
	}

	return s.sendOtpViaWhatsapp(phoneNumber, otp)
}

func (s *Service) VerifyOtp(phoneNumber, otp string) (*user.User, string, error) {
	userData, err := s.userRepo.FindByPhone(phoneNumber)
	if err == gorm.ErrRecordNotFound {
		return nil, "", fmt.Errorf("phone number not registered: %w", customerrors.ErrNotFound)
	}
	if err != nil {
		return nil, "", err
	}

	otpRecord, err := s.authRepo.GetPhoneOtp(phoneNumber)
	if err != nil {
		return nil, "", err
	}

	if !otpRecord.PhoneLoginOtp.Valid {
		return nil, "", fmt.Errorf("no verification code found: %w", customerrors.ErrNotFound)
	}

	err = bcrypt.CompareHashAndPassword([]byte(otpRecord.PhoneLoginOtp.String), []byte(otp))
	if err != nil {
		return nil, "", fmt.Errorf("incorrect verification code: %w", customerrors.ErrInvalid)
	}

	if otpRecord.PhoneLoginOtpExpires.Valid && 
		otpRecord.PhoneLoginOtpExpires.Time.Before(time.Now()) {
		return nil, "", fmt.Errorf("verification code expired: %w", customerrors.ErrExpired)
	}

	s.authRepo.ClearPhoneOtp(phoneNumber)

	token, err := jwt.GenerateJWT(userData.ID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return userData, token, nil
}


func (s *Service) generateOTP() (string, error) {
	max := big.NewInt(900000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()+100000), nil
}

func (s *Service) sendOtpViaWhatsapp(phone, otp string) error {
	log.Printf("Sending WhatsApp OTP to %s: %s", phone, otp)
	// TODO: Implement actual WhatsApp integration
	return nil
}