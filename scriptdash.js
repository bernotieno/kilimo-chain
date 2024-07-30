document.addEventListener('DOMContentLoaded', () => {
    const openSidebarBtn = document.getElementById('openSidebar');
    const closeSidebarBtn = document.getElementById('closeSidebar');
    const body = document.body;

    openSidebarBtn.addEventListener('click', () => {
        body.classList.add('sidebar-open');
    });

    closeSidebarBtn.addEventListener('click', () => {
        body.classList.remove('sidebar-open');
    });
});
