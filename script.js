function openSignUpForm() {
    closeForms();
    document.getElementById('signup-form').style.display = 'flex';
    document.body.style.filter = 'none';
}

function openSignInForm() {
    closeForms();
    document.getElementById('signin-form').style.display = 'flex';
    document.body.style.filter = 'none';
}

function openFarmerSignUpForm() {
    closeForms();
    document.getElementById('farmer-signup').style.display = 'flex';
    document.body.style.filter = 'none';
}

function closeForms() {
    document.getElementById('signup-form').style.display = 'none';
    document.getElementById('signin-form').style.display = 'none';
    document.getElementById('farmer-signup').style.display = 'none'
    document.body.style.filter = 'none';
}

document.addEventListener('DOMContentLoaded', () => {
    const banners = document.querySelectorAll('.banner');
    let delay = 0;

    banners.forEach(banner => {
        setTimeout(() => {
            banner.style.transform = 'translateX(0)';
        }, delay);
        delay += 300;
    });
});
