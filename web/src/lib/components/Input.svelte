<script lang="ts">
  export let type: string;
  export let name: string;
  export let label: string;
  export let required = false;
  export let placeholder = '';
</script>

<div class="input-wrapper">
  <label for={name} class="block mb-2 text-md font-bold dark:text-white">{label}</label>
  <input
    id={name}
    {type}
    {name}
    {required}
    {placeholder}
    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm focus:outline-zinc-500"
  />
</div>
