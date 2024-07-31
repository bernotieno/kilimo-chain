document.addEventListener('DOMContentLoaded', () => {
    const toggleBtn = document.querySelector('.toggle-btn');
    const sidebar = document.getElementById('sidebar');
    const closeBtn = document.getElementById('close-btn');
    const homeLink = document.getElementById('home-link');
    const infoLink = document.getElementById('info-link');
    const saleLink = document.getElementById('sale-link');
    const homeContent = document.getElementById('home');
    const infoContent = document.getElementById('info');
    const saleContent = document.getElementById('sale');

    toggleBtn.addEventListener('click', () => {
        sidebar.style.width = '250px';
    });

    closeBtn.addEventListener('click', () => {
        sidebar.style.width = '0';
    });

    homeLink.addEventListener('click', () => {
        setActiveContent(homeContent);
    const toggleBtn = document.querySelector('.toggle-btn');
    const sidebar = document.getElementById('sidebar');
    const closeBtn = document.getElementById('close-btn');
    const homeLink = document.getElementById('home-link');
    const infoLink = document.getElementById('info-link');
    const saleLink = document.getElementById('sale-link');
    const homeContent = document.getElementById('home');
    const infoContent = document.getElementById('info');
    const saleContent = document.getElementById('sale');

    toggleBtn.addEventListener('click', () => {
        sidebar.style.width = '250px';
    });

    closeBtn.addEventListener('click', () => {
        sidebar.style.width = '0';
    });

    homeLink.addEventListener('click', () => {
        setActiveContent(homeContent);
    });

    infoLink.addEventListener('click', () => {
        setActiveContent(infoContent);
    });

    saleLink.addEventListener('click', () => {
        setActiveContent(saleContent);
    });

    function setActiveContent(content) {
        [homeContent, infoContent, saleContent].forEach(c => c.classList.remove('active'));
        content.classList.add('active');
        sidebar.style.width = '0';
    }
    infoLink.addEventListener('click', () => {
        setActiveContent(infoContent);
    });

    saleLink.addEventListener('click', () => {
        setActiveContent(saleContent);
    });

    function setActiveContent(content) {
        [homeContent, infoContent, saleContent].forEach(c => c.classList.remove('active'));
        content.classList.add('active');
        sidebar.style.width = '0';
    }
});