import React, { FC, useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import { Alert } from "rsuite"
import Article from "../../models/Article"
import { requestGetArticle } from "../../apis/ArticleApi"
import ArticleTitle from "../articlePage/ArticleTitle"
import ArticleTags from "../articlePage/ArticleTags"
import ArticleLink from "../articlePage/ArticleLink"
import EditArticleContentMarkdown from "./EditArticleContentMarkdown"
import SimpleLayout from "../../component/layout/SimpleLayout"


const EditArticlePage: FC = () => {
  const {id} = useParams() as any
  const [loadFetching, setLoadFetching] = useState(false)
  const [article, setArticle] = useState(null as Article | null)

  useEffect(() => {
    setLoadFetching(true)
    requestGetArticle(id)
      .then(article => {
        setArticle(article)
      })
      .finally(() => setLoadFetching(false))
      .catch(err => Alert.error(err.toString()))
  }, [id])

  return (
    <SimpleLayout loading={loadFetching} size="lg">
      {renderArticle(article)}
    </SimpleLayout>
  )
}

const renderArticle = (article: Article | null) => {
  if (!article) {
    return null
  }

  return (
    <>
      <ArticleTitle article={article}/>
      <ArticleTags article={article}/>
      <ArticleLink article={article}/>
      <EditArticleContentMarkdown article={article}/>
    </>
  )
}

export default EditArticlePage
