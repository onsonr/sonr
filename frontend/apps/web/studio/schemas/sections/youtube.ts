import {RiYoutubeLine} from 'react-icons/ri';
import {defineField} from 'sanity';

const youtube = defineField({
	name: 'youtube',
	type: 'object',
	title: 'YouTube Embed',
	icon: RiYoutubeLine,
	fields: [
		{
			name: 'url',
			type: 'url',
			title: 'YouTube video URL',
			validation: (Rule) => Rule.required().uri({scheme: ['http']})
		},
		{
			name: 'autoPlay',
			type: 'boolean',
			title: 'Enable autoplay',
			validation: (Rule) => Rule.required()
		},
		{
			name: 'muted',
			type: 'boolean',
			title: 'Muted',
			validation: (Rule) => Rule.required()
		}
	]
});

export default youtube;
