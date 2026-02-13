# Exercise 2: The Payment Processor (Interfaces/ABCs)
# Scenario: Your app needs to handle multiple payment methods (Credit Card,
# PayPal, Crypto). You want to ensure that every payment method has a
# process_payment method.

# Task: Define an Abstract Base Class (Interface) called PaymentMethod.

# Requirement: Any subclass must implement process_payment(amount). Create two
# subclasses (CreditCard and PayPal) that print a unique message when the
# method is called.

# Interview Focus: This tests your understanding of Polymorphism and how to
# enforce "contracts" in your code to prevent errors.


from abc import ABC, abstractmethod


class PaymentMethod(ABC):

    @abstractmethod
    def process_payment(self, amount: float):
        pass


class CreditCard(PaymentMethod):
    def process_payment(self, amount: float):
        print(f"Processing credit card payment of ${amount:.2f}...")
        print("Validating CVV and expiry date...")