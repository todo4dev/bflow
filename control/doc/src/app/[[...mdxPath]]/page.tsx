import {generateStaticParamsFor, importPage} from 'nextra/pages'
import {useMDXComponents as getMDXComponents} from '../../mdx-components'

export const generateStaticParams = generateStaticParamsFor('mdxPath')

type PageProps = {
  params: Promise<{mdxPath?: string[]}>
  [key: string]: unknown
}

export async function generateMetadata({params}: PageProps) {
  const {mdxPath = []} = await params
  const {metadata} = await importPage(mdxPath)
  return metadata
}

const Wrapper = getMDXComponents().wrapper

export default async function Page({params, ...props}: PageProps) {
  const {mdxPath = []} = await params
  const {default: MDXContent, ...wrapperProps} = await importPage(mdxPath)
  return (
    <Wrapper {...wrapperProps}>
      <MDXContent {...props} params={{mdxPath}} />
    </Wrapper>
  )
}
