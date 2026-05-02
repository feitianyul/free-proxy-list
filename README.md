**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-05-02 19:53:05 UTC（2026-05-03 03:53:05 UTC+8）

**代理总数：44**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 44 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 44 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 149.51.42.10:3128 | ✓ 707ms | ✓ 1463ms | 否 | ✓ 1708ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1689ms | ✓ 1691ms | 否 | 否 | ✓ 1660ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1669ms | 否 | ✓ 1329ms | ✓ 1185ms | http |
| 45.167.124.71:999 | ✓ 1839ms | ✓ 1897ms | ✓ 1575ms | ✓ 1790ms | ✓ 1558ms | http |
| 218.108.131.186:17890 | ✓ 622ms | ✓ 767ms | ✓ 683ms | ✓ 879ms | ✓ 1854ms | http |
| 198.23.189.151:59394 | ✓ 854ms | ✓ 1460ms | ✓ 1067ms | ✓ 1535ms | ✓ 1322ms | http |
| 206.206.126.177:2412 | ✓ 1488ms | 否 | ✓ 861ms | ✓ 976ms | ✓ 940ms | http |
| 86.104.72.219:1081 | ✓ 611ms | ✓ 1369ms | ✓ 1479ms | ✓ 1455ms | 否 | http |
| 47.77.216.82:1080 | 否 | ✓ 825ms | ✓ 863ms | ✓ 767ms | ✓ 1638ms | http |
| 109.120.156.122:8090 | ✓ 1695ms | 否 | ✓ 1915ms | 否 | ✓ 1784ms | http |
| 117.236.124.166:3128 | ✓ 1565ms | 否 | ✓ 1996ms | 否 | ✓ 1589ms | http |
| 59.46.216.131:30001 | ✓ 1823ms | 否 | ✓ 879ms | ✓ 1073ms | ✓ 886ms | http |
| 47.85.51.197:1080 | ✓ 282ms | ✓ 1518ms | ✓ 1197ms | 否 | ✓ 899ms | http |
| 120.92.108.86:7890 | ✓ 980ms | 否 | ✓ 1071ms | 否 | ✓ 1616ms | http |
| 80.92.204.47:1081 | ✓ 1227ms | ✓ 1473ms | ✓ 1484ms | 否 | 否 | http |
| 103.172.70.173:8080 | 否 | ✓ 1854ms | ✓ 1536ms | ✓ 1337ms | 否 | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1522ms | ✓ 1565ms | ✓ 1289ms | http |
| 86.104.72.219:1082 | ✓ 430ms | ✓ 1115ms | 否 | ✓ 1453ms | 否 | http |
| 149.51.42.10:8080 | ✓ 1733ms | ✓ 1837ms | 否 | ✓ 1510ms | 否 | http |
| 72.11.150.178:6005 | ✓ 1710ms | ✓ 1367ms | ✓ 1052ms | ✓ 1444ms | ✓ 1885ms | http |
| 45.174.77.224:999 | 否 | 否 | ✓ 1043ms | ✓ 1158ms | ✓ 1042ms | http |
| 34.96.238.40:8080 | ✓ 1275ms | ✓ 1034ms | ✓ 1521ms | ✓ 959ms | 否 | http |
| 92.119.127.208:6005 | ✓ 1119ms | ✓ 1865ms | ✓ 1435ms | 否 | ✓ 1892ms | http |
| 180.103.19.166:1080 | ✓ 1221ms | ✓ 1110ms | ✓ 1389ms | ✓ 1281ms | ✓ 1078ms | http |
| 152.32.132.190:7890 | ✓ 856ms | ✓ 1599ms | 否 | ✓ 1148ms | 否 | http |
| 223.84.151.86:30005 | 否 | ✓ 1829ms | ✓ 1977ms | ✓ 1556ms | ✓ 1587ms | http |
| 8.154.21.175:3128 | ✓ 708ms | 否 | ✓ 674ms | ✓ 924ms | ✓ 1269ms | http |
| 152.42.177.32:8888 | ✓ 1338ms | 否 | ✓ 1252ms | ✓ 1270ms | 否 | http |
| 154.90.48.209:9090 | ✓ 1554ms | 否 | ✓ 786ms | ✓ 1447ms | ✓ 905ms | http |
| 47.112.25.109:7890 | 否 | ✓ 1760ms | ✓ 852ms | 否 | ✓ 1768ms | http |
| 190.12.150.244:999 | ✓ 1047ms | ✓ 1915ms | ✓ 1416ms | 否 | 否 | http |
| 116.171.106.111:3443 | ✓ 1223ms | ✓ 1216ms | ✓ 1196ms | ✓ 1518ms | ✓ 1391ms | http |
| 121.147.253.205:3164 | 否 | 否 | ✓ 870ms | ✓ 976ms | ✓ 943ms | http |
| 222.102.86.137:3040 | 否 | 否 | ✓ 866ms | ✓ 959ms | ✓ 1142ms | http |
| 3.101.133.120:80 | 否 | 否 | ✓ 1260ms | ✓ 1160ms | ✓ 672ms | http |
| 103.157.200.126:3128 | 否 | 否 | ✓ 1845ms | ✓ 1973ms | ✓ 1551ms | http |
| 116.171.106.78:3443 | 否 | ✓ 1375ms | ✓ 1626ms | ✓ 1619ms | ✓ 1646ms | http |
| 101.32.243.189:80 | ✓ 1283ms | 否 | 否 | ✓ 1521ms | ✓ 1383ms | http |
| 121.230.8.55:1080 | ✓ 921ms | ✓ 1163ms | ✓ 1092ms | ✓ 1469ms | ✓ 1054ms | http |
| 61.52.131.172:8443 | ✓ 1294ms | ✓ 1947ms | ✓ 1346ms | ✓ 1978ms | ✓ 769ms | http |
| 116.171.106.26:3443 | ✓ 1913ms | 否 | ✓ 1409ms | ✓ 1714ms | 否 | http |
| 121.230.8.136:1080 | ✓ 1783ms | ✓ 1727ms | ✓ 1334ms | 否 | 否 | http |
| 103.39.51.207:8080 | ✓ 1336ms | 否 | 否 | ✓ 1993ms | ✓ 1446ms | http |
| 120.92.212.16:7890 | ✓ 1329ms | ✓ 1978ms | ✓ 818ms | 否 | 否 | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
