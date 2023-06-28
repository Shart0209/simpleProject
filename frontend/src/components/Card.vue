<template>
    <div class="card">
        <h4 class="card-header text-center">
            <slot name="header"></slot>
        </h4>
        <div class="card-body">
            <div class="card-text fs-5 text">
                <slot name="body"></slot>
            </div>
        </div>
        <div class="card-footer text-muted">
            <slot name="footer"></slot>
        </div>
    </div>
</template>