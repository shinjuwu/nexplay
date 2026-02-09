<template>
  <Transition
    name="slide"
    enter-active-class="slide-enter-active"
    leave-active-class="slide-leave-active"
    @before-enter="onBeforeEnter"
    @enter="onEnter"
    @after-enter="onAfterEnter"
    @before-leave="onBeforeLeave"
    @leave="onLeave"
    @after-leave="onAfterLeave"
  >
    <slot></slot>
  </Transition>
</template>

<script setup>
function onBeforeEnter(element) {
  requestAnimationFrame(() => {
    if (!element.style.height) {
      element.style.height = '0px'
    }
    element.style.display = null
  })
}

function onEnter(element) {
  requestAnimationFrame(() => {
    requestAnimationFrame(() => {
      element.style.height = `${element.scrollHeight}px`
    })
  })
}

function onAfterEnter(element) {
  element.style.height = null
}

function onBeforeLeave(element) {
  requestAnimationFrame(() => {
    if (!element.style.height) {
      element.style.height = `${element.offsetHeight}px`
    }
  })
}

function onLeave(element) {
  requestAnimationFrame(() => {
    requestAnimationFrame(() => {
      element.style.height = '0px'
    })
  })
}

function onAfterLeave(element) {
  element.style.height = null
}
</script>

<style scoped>
.slide-enter-active,
.slide-leave-active {
  overflow: hidden;
  transition: height 0.25s linear;
}
</style>
