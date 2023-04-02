<script setup>

const props = defineProps({
  id: String,
  close_func: Function,
})

</script>

<template>
  <div 
    :id="props.id" 
    class="modal fade" 
    data-bs-backdrop="static" 
    data-bs-keyboard="false" 
    tabindex="-1"
    :aria-labelledby="props.id + '_label'" 
    aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 
            class="modal-title fs-5"
            :id="props.id + 'label'"
            >
            <slot name="title"></slot>
          </h5>
          <button 
          type="button" 
          class="btn-close" 
          @click="props.close_func"
          >
          </button>
        </div>
        <div class="modal-body">
          <slot name="body"></slot>
        </div>
        <div class="modal-footer">
          <slot name="footer"></slot>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped></style>