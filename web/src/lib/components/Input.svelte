<script lang="ts">
  export let type: string;
  export let name: string;
  export let label: string;
  export let required = false;
  export let placeholder = '';
  export let value = '';
  export let subLabel = ''
</script>

<div class="input-wrapper">
  <label for={name} class="block mb-2 text-md font-bold dark:text-white">
    {label}
    {#if required}
      <span>*</span>
    {/if}
    {#if subLabel}
      <span class="sub-label font-normal italic text-xs">({subLabel})</span>
    {/if}
  </label>
  <input
    id={name}
    {type}
    {name}
    {required}
    {placeholder}
    {value}
    class="flex h-10 w-full rounded-md border border-input bg-background px-3
    py-2 text-sm outline-none focus:outline-zinc-500 dark:bg-zinc-800 dark:border-zinc-700
    dark:focus:outline-zinc-600"
  />
</div>
