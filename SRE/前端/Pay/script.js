document.addEventListener('DOMContentLoaded', function() {
    // 支付方式选择
    const creditCardOption = document.getElementById('credit-card-option');
    const paypalOption = document.getElementById('paypal-option');
    const creditCardForm = document.getElementById('credit-card-form');
    const paypalForm = document.getElementById('paypal-form');
    
    creditCardOption.addEventListener('click', function() {
        creditCardOption.classList.add('active');
        paypalOption.classList.remove('active');
        creditCardForm.classList.remove('hidden');
        paypalForm.classList.add('hidden');
    });
    
    paypalOption.addEventListener('click', function() {
        paypalOption.classList.add('active');
        creditCardOption.classList.remove('active');
        paypalForm.classList.remove('hidden');
        creditCardForm.classList.add('hidden');
    });
    
    // 信用卡号格式化
    const cardNumberInput = document.getElementById('card-number');
    cardNumberInput.addEventListener('input', function(e) {
        let value = e.target.value.replace(/\D/g, '');
        let formattedValue = '';
        
        for (let i = 0; i < value.length; i++) {
            if (i > 0 && i % 4 === 0) {
                formattedValue += ' ';
            }
            formattedValue += value[i];
        }
        
        e.target.value = formattedValue;
    });
    
    // 有效期格式化
    const expirationInput = document.getElementById('expiration');
    expirationInput.addEventListener('input', function(e) {
        let value = e.target.value.replace(/\D/g, '');
        let formattedValue = '';
        
        if (value.length > 0) {
            formattedValue = value.substring(0, 2);
            
            if (value.length > 2) {
                formattedValue += ' / ' + value.substring(2, 4);
            }
        }
        
        e.target.value = formattedValue;
    });
    
    // CVC只允许输入数字
    const cvcInput = document.getElementById('cvc');
    cvcInput.addEventListener('input', function(e) {
        e.target.value = e.target.value.replace(/\D/g, '');
    });
    
    // 表单提交
    const checkoutButton = document.querySelector('.checkout-button');
    checkoutButton.addEventListener('click', function() {
        // 这里可以添加表单验证逻辑
        if (creditCardForm.classList.contains('hidden')) {
            alert('PayPal 支付处理中...');
        } else {
            const cardNumber = cardNumberInput.value.replace(/\s/g, '');
            const expiration = expirationInput.value;
            const cvc = cvcInput.value;
            
            if (!cardNumber || cardNumber.length < 13) {
                alert('请输入有效的卡号');
                return;
            }
            
            if (!expiration || expiration.length < 7) {
                alert('请输入有效的到期日期');
                return;
            }
            
            if (!cvc || cvc.length < 3) {
                alert('请输入有效的CVC码');
                return;
            }
            
            alert('信用卡支付处理中...');
        }
    });
}); 