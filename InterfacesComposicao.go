package main

import (
	"errors"
	"fmt"
	"sync"
)

// -> Notifier define para envio de notificações
type Notifier interface {
	Send(message string) error
	Name() string
}

// Implementação
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) error {
	if message == "" {
		return errors.New("mensagem vazia para Email")
	}
	fmt.Println("Email enviado:", message)
	return nil
}

func (e *EmailNotifier) Name() string {
	return "email"
}

type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) error {
	if message == "" {
		return errors.New("mensagem vazia para SMS")
	}
	fmt.Println("SMS enviado:", message)
	return nil
}

func (s *SMSNotifier) Name() string {
	return "sms"
}

type PushNotifier struct{}

func (p *PushNotifier) Send(message string) error {
	if message == "" {
		return errors.New("mensagem vazia para Push")
	}
	fmt.Println("Push enviado:", message)
	return nil
}

func (p *PushNotifier) Name() string {
	return "push"
}

// Serviço de Notificação

type NotificationService struct {
	notifiers map[string]Notifier
	mu        sync.Mutex
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		notifiers: make(map[string]Notifier),
	}
}

// AddNotifier adiciona um notificador dinamicamente
func (s *NotificationService) AddNotifier(n Notifier) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.notifiers[n.Name()] = n
}

// RemoveNotifier remove um notificador pelo nome
func (s *NotificationService) RemoveNotifier(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.notifiers, name)
}

// Broadcast envia a mensagem para todos os notificadores
func (s *NotificationService) Broadcast(message string) []error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var errs []error

	for _, notifier := range s.notifiers {
		if err := notifier.Send(message); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
