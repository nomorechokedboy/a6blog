<script setup lang="ts">
import { onClickOutside } from '@vueuse/core'
import { ActionIcon } from 'ui-vue'

async function logout() {
	removeToken()
	notifyController?.abort()
	await navigateTo('/login')
}

function handleToggle() {
	toggle.value = !toggle.value
}

const appConfig = useRuntimeConfig()
const { data: userProfile, isLoading: isUserProfileLoading } = useUserProfile()
const toggle = ref(false)
const userAvatar = computed(() =>
	userProfile.value?.avatar
		? `${appConfig.public.mediaUrl}${userProfile.value?.avatar}`
		: `${appConfig.public.dicebearMedia}${userProfile.value?.name}`
)
const target = ref<HTMLDivElement | null>(null)
const notifyController = inject<AbortController>('notifyController')

onClickOutside(target, () => (toggle.value = false))
</script>

<template>
	<div class="relative inline-block" ref="target">
		<ActionIcon
			radius="xl"
			size="lg"
			variant="subtle"
			@click="handleToggle"
			class="focus:ring-4 focus:outline-none focus:ring-gray-300"
		>
			<Avatar
				v-if="!isUserProfileLoading"
				width="32"
				:src="userAvatar"
			/>
			<div class="skeleton !rounded-full" v-else>
				<div class="w-6 h-6 md:!w-8 md:h-8" />
			</div>
		</ActionIcon>
		<Dropdown :open="toggle">
			<div
				class="pb-2 mb-2 hover:underline hover:rounded-lg px-4 py-3 hover:bg-gray-100"
			>
				<NuxtLink
					:to="`/${
						userProfile?.id || 'User name'
					}`"
				>
					<span class="fw-medium block">
						{{ userProfile?.name || 'Oof' }}
					</span>
					<small class="text-gray-400"
						>@{{
							userProfile?.username ||
							'Oof'
						}}
					</small>
				</NuxtLink>
			</div>
			<ul
				class="py-2 text-gray-700"
				aria-labelledby="dropdownDefaultButton"
			>
				<!-- <li>
                <a href="#"
                    class="block px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Dashboard</a>
            </li> -->
				<li>
					<NuxtLink
						to="/new"
						class="block px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100"
					>
						Create Post</NuxtLink
					>
				</li>
				<!-- <li>
                <a href="#"
                    class="block px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Reading
                    list</a>
            </li>
            <li>
                <a href="#"
                    class="block px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Settings</a>
            </li> -->
			</ul>
			<div class="pt-2">
				<div
					class="cursor-pointer px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100"
					@click="logout"
				>
					Sign out
				</div>
			</div>
		</Dropdown>
	</div>
</template>
