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

最后更新：2026-03-01 09:35:11 UTC（2026-03-01 17:35:11 UTC+8）

**代理总数：60**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 36.147.78.166:80 | 否 | ✓ 1541ms | ✓ 1652ms | ✓ 1878ms | 否 | http |
| 148.135.85.87:1080 | ✓ 1518ms | ✓ 1605ms | 否 | 否 | ✓ 914ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 710ms | ✓ 962ms | ✓ 1920ms | http |
| 59.46.216.131:30001 | ✓ 832ms | ✓ 1014ms | ✓ 953ms | ✓ 1113ms | ✓ 932ms | http |
| 115.231.181.40:8128 | ✓ 1064ms | ✓ 1829ms | ✓ 1253ms | ✓ 1104ms | ✓ 844ms | http |
| 195.123.209.48:3128 | ✓ 1763ms | ✓ 1742ms | ✓ 1435ms | 否 | 否 | http |
| 103.236.64.247:8888 | ✓ 1213ms | ✓ 1863ms | 否 | 否 | ✓ 704ms | http |
| 74.208.234.198:443 | ✓ 979ms | ✓ 1877ms | 否 | ✓ 1260ms | ✓ 1362ms | http |
| 14.56.107.244:3128 | ✓ 1679ms | 否 | ✓ 764ms | ✓ 1302ms | ✓ 1222ms | http |
| 3.213.157.4:3128 | ✓ 1099ms | 否 | ✓ 1252ms | ✓ 1654ms | ✓ 1431ms | http |
| 107.173.83.243:8002 | ✓ 786ms | ✓ 745ms | ✓ 817ms | ✓ 1035ms | 否 | http |
| 81.70.169.194:80 | ✓ 825ms | ✓ 1562ms | ✓ 800ms | ✓ 1208ms | ✓ 1829ms | http |
| 101.43.255.96:80 | ✓ 768ms | ✓ 1640ms | 否 | 否 | ✓ 1947ms | http |
| 205.209.118.30:3138 | ✓ 1062ms | 否 | ✓ 302ms | ✓ 1392ms | ✓ 1035ms | http |
| 107.174.133.10:3128 | 否 | 否 | ✓ 207ms | ✓ 1704ms | ✓ 537ms | http |
| 35.225.22.61:80 | ✓ 386ms | ✓ 1369ms | ✓ 1053ms | ✓ 1372ms | ✓ 1253ms | http |
| 103.84.95.54:7890 | ✓ 616ms | 否 | ✓ 713ms | ✓ 819ms | 否 | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1605ms | ✓ 1103ms | ✓ 1025ms | http |
| 120.92.212.16:7890 | ✓ 1330ms | 否 | ✓ 1125ms | ✓ 1168ms | 否 | http |
| 147.45.216.148:1080 | ✓ 1590ms | 否 | ✓ 1475ms | ✓ 1920ms | ✓ 1395ms | http |
| 35.234.17.221:8080 | 否 | ✓ 1676ms | 否 | ✓ 958ms | ✓ 833ms | http |
| 121.230.8.211:1080 | ✓ 926ms | ✓ 1123ms | ✓ 1122ms | 否 | 否 | http |
| 158.178.237.243:3128 | ✓ 1395ms | 否 | ✓ 1877ms | ✓ 1299ms | ✓ 1519ms | http |
| 121.230.8.109:1080 | ✓ 959ms | ✓ 1073ms | ✓ 942ms | ✓ 1800ms | ✓ 1343ms | http |
| 146.56.182.165:3128 | ✓ 1060ms | 否 | 否 | ✓ 1505ms | ✓ 1119ms | http |
| 42.115.247.250:10031 | ✓ 1847ms | 否 | ✓ 1690ms | 否 | ✓ 1810ms | http |
| 1.12.62.237:8080 | ✓ 1825ms | ✓ 1383ms | 否 | ✓ 1302ms | ✓ 1185ms | http |
| 188.166.208.168:9876 | ✓ 1441ms | 否 | ✓ 1000ms | ✓ 1010ms | ✓ 855ms | http |
| 36.147.78.166:443 | ✓ 1947ms | 否 | ✓ 1671ms | ✓ 1797ms | ✓ 1582ms | http |
| 168.235.110.63:3128 | 否 | 否 | ✓ 1450ms | ✓ 1541ms | ✓ 1396ms | http |
| 61.72.110.24:3128 | ✓ 1467ms | 否 | ✓ 1414ms | ✓ 1976ms | ✓ 801ms | http |
| 61.72.110.94:3128 | ✓ 1677ms | 否 | ✓ 1565ms | 否 | ✓ 1427ms | http |
| 14.56.118.164:3128 | ✓ 932ms | ✓ 1711ms | 否 | 否 | ✓ 824ms | http |
| 103.215.36.88:15088 | ✓ 863ms | ✓ 1218ms | ✓ 937ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 845ms | ✓ 1776ms | ✓ 941ms | ✓ 1658ms | ✓ 1368ms | http |
| 52.188.28.218:3128 | ✓ 1427ms | 否 | ✓ 466ms | 否 | ✓ 1300ms | http |
| 23.224.193.43:3128 | ✓ 842ms | 否 | 否 | ✓ 1262ms | ✓ 963ms | http |
| 23.224.193.44:3128 | ✓ 766ms | 否 | ✓ 1707ms | ✓ 1261ms | ✓ 1617ms | http |
| 23.224.193.45:3128 | ✓ 675ms | 否 | ✓ 1694ms | ✓ 1236ms | ✓ 1677ms | http |
| 45.140.147.155:1081 | ✓ 710ms | 否 | ✓ 1154ms | ✓ 1805ms | ✓ 1256ms | http |
| 45.140.147.155:1082 | ✓ 586ms | 否 | ✓ 674ms | 否 | ✓ 1127ms | http |
| 62.113.119.14:8080 | ✓ 781ms | 否 | ✓ 1186ms | 否 | ✓ 1432ms | http |
| 160.250.5.22:1 | ✓ 1638ms | 否 | ✓ 1510ms | ✓ 1177ms | ✓ 923ms | http |
| 167.160.184.231:6005 | ✓ 1189ms | 否 | ✓ 1123ms | 否 | ✓ 1113ms | http |
| 23.224.193.42:3128 | ✓ 797ms | ✓ 1505ms | 否 | 否 | ✓ 980ms | http |
| 160.250.4.245:1 | ✓ 1628ms | 否 | ✓ 1900ms | ✓ 1185ms | ✓ 941ms | http |
| 165.227.5.10:8888 | ✓ 1635ms | ✓ 1300ms | ✓ 734ms | ✓ 1924ms | ✓ 1514ms | http |
| 103.39.51.190:8080 | ✓ 1641ms | 否 | 否 | ✓ 1235ms | ✓ 1274ms | http |
| 177.243.209.133:999 | ✓ 884ms | 否 | ✓ 1121ms | ✓ 1179ms | ✓ 988ms | http |
| 61.72.110.54:3128 | ✓ 1510ms | 否 | 否 | ✓ 1901ms | ✓ 1156ms | http |
| 14.56.118.244:3128 | ✓ 1383ms | 否 | ✓ 1799ms | 否 | ✓ 1945ms | http |
| 45.140.147.82:1081 | ✓ 598ms | ✓ 1408ms | 否 | ✓ 1927ms | 否 | http |
| 14.56.118.214:3128 | ✓ 1670ms | ✓ 1571ms | ✓ 1428ms | 否 | 否 | http |
| 180.127.149.244:1080 | ✓ 968ms | 否 | ✓ 758ms | ✓ 1118ms | ✓ 1223ms | http |
| 120.55.163.237:10086 | ✓ 743ms | ✓ 837ms | ✓ 1264ms | ✓ 1713ms | ✓ 955ms | http |
| 103.131.19.42:8181 | ✓ 1444ms | 否 | ✓ 1384ms | ✓ 1406ms | ✓ 1455ms | http |
| 103.104.99.29:80 | ✓ 1320ms | 否 | ✓ 1047ms | ✓ 1434ms | 否 | http |
| 103.104.99.89:80 | ✓ 1975ms | 否 | ✓ 1563ms | ✓ 1464ms | ✓ 1495ms | http |
| 207.254.71.62:8088 | ✓ 1153ms | 否 | ✓ 1674ms | 否 | ✓ 1860ms | http |
| 142.171.85.32:1080 | ✓ 373ms | ✓ 929ms | ✓ 981ms | ✓ 895ms | ✓ 993ms | http |

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
