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

最后更新：2026-03-04 08:45:28 UTC（2026-03-04 16:45:28 UTC+8）

**代理总数：51**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 51 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 51 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.225.116.115:1080 | ✓ 1071ms | ✓ 1598ms | ✓ 1137ms | ✓ 1025ms | ✓ 895ms | http |
| 95.85.252.153:21064 | ✓ 1185ms | ✓ 1649ms | ✓ 1672ms | 否 | 否 | http |
| 125.128.12.14:3128 | ✓ 613ms | ✓ 1047ms | 否 | ✓ 910ms | ✓ 725ms | http |
| 115.231.181.40:8128 | ✓ 1153ms | 否 | ✓ 1783ms | 否 | ✓ 1092ms | http |
| 59.46.216.131:30001 | ✓ 916ms | 否 | 否 | ✓ 1793ms | ✓ 1417ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 820ms | ✓ 1347ms | ✓ 935ms | http |
| 61.72.221.194:3128 | ✓ 1678ms | 否 | ✓ 1424ms | ✓ 1032ms | ✓ 785ms | http |
| 120.92.212.16:7890 | ✓ 920ms | ✓ 1369ms | ✓ 937ms | ✓ 1218ms | ✓ 919ms | http |
| 81.70.169.194:80 | ✓ 986ms | ✓ 1236ms | ✓ 963ms | ✓ 1286ms | ✓ 967ms | http |
| 101.43.255.96:80 | ✓ 987ms | 否 | ✓ 920ms | ✓ 1524ms | ✓ 1947ms | http |
| 120.92.212.16:8890 | ✓ 913ms | ✓ 1573ms | ✓ 1165ms | ✓ 1361ms | 否 | http |
| 61.72.221.234:3128 | ✓ 882ms | ✓ 1566ms | ✓ 1745ms | 否 | 否 | http |
| 91.193.240.157:9877 | ✓ 1169ms | 否 | ✓ 1314ms | 否 | ✓ 1869ms | http |
| 111.79.111.126:3128 | 否 | 否 | ✓ 1231ms | ✓ 1868ms | ✓ 1571ms | http |
| 35.234.17.221:8080 | ✓ 787ms | ✓ 1573ms | ✓ 855ms | 否 | ✓ 914ms | http |
| 122.2.48.121:8080 | ✓ 1090ms | 否 | 否 | ✓ 1194ms | ✓ 1341ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1013ms | ✓ 786ms | ✓ 761ms | http |
| 45.140.147.155:1081 | ✓ 591ms | 否 | 否 | ✓ 1406ms | ✓ 1684ms | http |
| 45.140.147.155:1082 | ✓ 595ms | ✓ 1441ms | 否 | ✓ 1967ms | ✓ 1642ms | http |
| 121.141.161.55:1080 | 否 | ✓ 1569ms | 否 | ✓ 910ms | ✓ 1778ms | http |
| 121.230.8.245:1080 | 否 | 否 | ✓ 1871ms | ✓ 1695ms | ✓ 979ms | http |
| 58.220.95.12:12120 | ✓ 1492ms | 否 | ✓ 1873ms | 否 | ✓ 1967ms | http |
| 64.181.240.152:3128 | ✓ 552ms | 否 | ✓ 698ms | ✓ 1394ms | ✓ 744ms | http |
| 192.166.82.55:1080 | ✓ 1952ms | 否 | ✓ 1368ms | ✓ 1941ms | ✓ 1674ms | http |
| 150.241.77.125:3128 | ✓ 1232ms | 否 | ✓ 603ms | ✓ 1735ms | 否 | http |
| 121.230.9.115:1080 | ✓ 1217ms | ✓ 1213ms | ✓ 1084ms | ✓ 1423ms | ✓ 916ms | http |
| 58.220.95.12:11904 | ✓ 1744ms | 否 | ✓ 1300ms | ✓ 1158ms | ✓ 921ms | http |
| 61.52.131.172:8443 | ✓ 1241ms | ✓ 1057ms | ✓ 827ms | ✓ 1074ms | ✓ 864ms | http |
| 205.209.118.30:3138 | ✓ 967ms | 否 | 否 | ✓ 1362ms | ✓ 1076ms | http |
| 45.136.198.40:3128 | ✓ 1125ms | 否 | ✓ 1748ms | 否 | ✓ 1928ms | http |
| 5.101.0.233:3128 | 否 | 否 | ✓ 1694ms | ✓ 1884ms | ✓ 1832ms | http |
| 103.215.36.88:15968 | 否 | ✓ 1950ms | 否 | ✓ 1539ms | ✓ 1204ms | http |
| 103.215.36.88:16988 | ✓ 1214ms | ✓ 1207ms | ✓ 1076ms | ✓ 1296ms | ✓ 1570ms | http |
| 180.103.19.49:1080 | ✓ 1126ms | ✓ 1646ms | 否 | ✓ 1681ms | ✓ 1694ms | http |
| 77.83.203.6:443 | ✓ 1126ms | ✓ 1969ms | ✓ 1478ms | ✓ 1957ms | ✓ 1169ms | http |
| 94.176.3.43:7443 | ✓ 937ms | 否 | ✓ 1442ms | 否 | ✓ 1428ms | http |
| 103.215.36.88:10864 | ✓ 977ms | ✓ 1262ms | ✓ 1619ms | 否 | ✓ 1526ms | http |
| 46.249.103.192:443 | ✓ 759ms | 否 | ✓ 1504ms | ✓ 1905ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1682ms | 否 | 否 | ✓ 1444ms | ✓ 1376ms | http |
| 14.56.107.244:3128 | ✓ 933ms | 否 | ✓ 1024ms | ✓ 1202ms | ✓ 752ms | http |
| 62.113.119.14:8080 | ✓ 921ms | 否 | ✓ 1463ms | ✓ 1954ms | ✓ 1498ms | http |
| 167.172.253.162:4857 | ✓ 1684ms | ✓ 1523ms | 否 | ✓ 1190ms | 否 | http |
| 121.230.8.74:1080 | ✓ 956ms | 否 | ✓ 1156ms | 否 | ✓ 962ms | http |
| 103.112.163.46:8080 | 否 | 否 | ✓ 1651ms | ✓ 1637ms | ✓ 1589ms | http |
| 122.52.16.235:8082 | 否 | 否 | ✓ 1527ms | ✓ 1275ms | ✓ 1476ms | http |
| 103.106.216.133:8097 | 否 | 否 | ✓ 1965ms | ✓ 1531ms | ✓ 1427ms | http |
| 91.233.223.147:3128 | ✓ 1571ms | 否 | ✓ 1518ms | 否 | ✓ 1777ms | http |
| 90.84.188.97:8000 | ✓ 1155ms | 否 | ✓ 729ms | 否 | ✓ 1892ms | http |
| 103.215.36.88:17090 | ✓ 991ms | ✓ 1972ms | ✓ 1138ms | ✓ 1950ms | ✓ 940ms | http |
| 121.230.8.97:1080 | ✓ 1163ms | ✓ 1824ms | ✓ 1234ms | 否 | 否 | http |
| 202.129.206.239:3128 | ✓ 1509ms | 否 | ✓ 1456ms | 否 | ✓ 1425ms | http |

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
