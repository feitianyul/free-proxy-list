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

最后更新：2026-03-04 05:40:22 UTC（2026-03-04 13:40:22 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 3.213.157.4:3128 | ✓ 451ms | 否 | ✓ 962ms | ✓ 1718ms | ✓ 1307ms | http |
| 95.85.252.153:21064 | ✓ 1339ms | ✓ 1842ms | ✓ 1558ms | 否 | 否 | http |
| 121.141.161.55:1080 | ✓ 856ms | 否 | ✓ 991ms | ✓ 1005ms | ✓ 829ms | http |
| 61.72.221.234:3128 | ✓ 641ms | 否 | 否 | ✓ 1442ms | ✓ 955ms | http |
| 211.171.114.154:3128 | ✓ 1195ms | ✓ 1231ms | 否 | 否 | ✓ 1335ms | http |
| 45.140.147.155:1081 | ✓ 622ms | 否 | ✓ 1800ms | ✓ 1896ms | 否 | http |
| 4.216.195.194:3128 | ✓ 552ms | ✓ 1674ms | ✓ 1713ms | 否 | ✓ 775ms | http |
| 115.231.181.40:8128 | ✓ 1618ms | ✓ 1040ms | ✓ 1253ms | 否 | ✓ 871ms | http |
| 59.46.216.131:30001 | ✓ 1016ms | ✓ 1345ms | ✓ 1137ms | ✓ 1432ms | 否 | http |
| 8.219.97.248:80 | ✓ 1743ms | 否 | ✓ 1958ms | ✓ 1629ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1543ms | 否 | ✓ 1468ms | 否 | ✓ 1397ms | http |
| 35.225.22.61:80 | ✓ 934ms | 否 | ✓ 540ms | ✓ 1041ms | ✓ 1133ms | http |
| 159.89.31.62:8080 | ✓ 657ms | 否 | ✓ 1280ms | 否 | ✓ 1394ms | http |
| 120.92.212.16:8890 | ✓ 939ms | ✓ 1208ms | ✓ 1795ms | ✓ 1451ms | ✓ 910ms | http |
| 103.86.131.62:80 | ✓ 1298ms | 否 | ✓ 1331ms | ✓ 1223ms | ✓ 1010ms | http |
| 91.193.240.157:9877 | ✓ 874ms | 否 | ✓ 904ms | 否 | ✓ 1592ms | http |
| 70.22.175.232:3128 | 否 | ✓ 1222ms | ✓ 361ms | ✓ 1283ms | ✓ 988ms | http |
| 160.250.4.245:1 | ✓ 1688ms | 否 | ✓ 1518ms | 否 | ✓ 969ms | http |
| 160.250.5.22:1 | ✓ 1697ms | 否 | 否 | ✓ 1288ms | ✓ 1327ms | http |
| 121.230.9.168:1080 | 否 | ✓ 1881ms | ✓ 1282ms | ✓ 1399ms | ✓ 1233ms | http |
| 120.92.212.16:7890 | ✓ 931ms | ✓ 1435ms | ✓ 1552ms | 否 | ✓ 1427ms | http |
| 160.250.4.13:1 | ✓ 1716ms | 否 | 否 | ✓ 1768ms | ✓ 1267ms | http |
| 121.204.158.249:3128 | ✓ 1177ms | ✓ 1417ms | ✓ 1676ms | 否 | ✓ 1520ms | http |
| 205.209.118.30:3138 | ✓ 343ms | ✓ 1279ms | ✓ 328ms | 否 | 否 | http |
| 122.2.48.121:8080 | ✓ 1197ms | 否 | ✓ 1172ms | ✓ 1247ms | ✓ 1230ms | http |
| 37.27.100.80:443 | 否 | 否 | ✓ 1790ms | ✓ 1767ms | ✓ 1504ms | http |
| 125.128.12.14:3128 | ✓ 1495ms | ✓ 1799ms | ✓ 1094ms | ✓ 1914ms | ✓ 1511ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1206ms | 否 | ✓ 1341ms | ✓ 989ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1546ms | ✓ 1321ms | ✓ 1962ms | http |
| 103.84.95.54:7890 | ✓ 1020ms | 否 | 否 | ✓ 969ms | ✓ 630ms | http |
| 180.127.149.247:1080 | ✓ 1620ms | ✓ 1229ms | ✓ 1957ms | 否 | 否 | http |
| 101.43.255.96:80 | ✓ 1523ms | ✓ 1278ms | 否 | 否 | ✓ 1712ms | http |
| 121.230.8.214:1080 | ✓ 1165ms | ✓ 1641ms | ✓ 1485ms | ✓ 1629ms | ✓ 1250ms | http |
| 81.70.169.194:80 | ✓ 1658ms | ✓ 1484ms | ✓ 947ms | ✓ 1231ms | ✓ 974ms | http |
| 121.230.8.97:1080 | ✓ 1161ms | 否 | ✓ 1062ms | ✓ 1818ms | ✓ 1127ms | http |
| 1.12.62.237:8080 | ✓ 1670ms | 否 | ✓ 1464ms | ✓ 1673ms | ✓ 1846ms | http |
| 5.75.196.26:40000 | ✓ 1199ms | 否 | ✓ 1395ms | ✓ 1573ms | ✓ 1388ms | http |
| 45.136.198.40:3128 | ✓ 1258ms | ✓ 1764ms | ✓ 933ms | ✓ 1732ms | ✓ 1300ms | http |
| 103.166.185.54:3128 | 否 | ✓ 1614ms | ✓ 1356ms | ✓ 1177ms | ✓ 1040ms | http |
| 61.72.221.194:3128 | ✓ 1328ms | 否 | ✓ 1118ms | 否 | ✓ 1554ms | http |
| 154.90.48.209:9090 | ✓ 1864ms | 否 | ✓ 1432ms | 否 | ✓ 1752ms | http |
| 66.228.47.125:110 | ✓ 1407ms | 否 | ✓ 1487ms | 否 | ✓ 1986ms | http |
| 113.176.92.71:3128 | ✓ 1762ms | ✓ 1534ms | ✓ 1469ms | 否 | ✓ 1324ms | http |
| 83.219.250.8:62920 | ✓ 1250ms | ✓ 1844ms | 否 | 否 | ✓ 1941ms | http |
| 103.215.36.88:13763 | ✓ 1058ms | ✓ 1521ms | ✓ 1573ms | ✓ 1572ms | ✓ 1321ms | http |
| 35.234.17.221:8080 | ✓ 839ms | 否 | 否 | ✓ 947ms | ✓ 947ms | http |
| 150.249.255.91:3128 | ✓ 712ms | ✓ 1400ms | ✓ 514ms | ✓ 799ms | ✓ 649ms | http |
| 221.202.27.194:10809 | ✓ 1972ms | 否 | 否 | ✓ 1703ms | ✓ 1787ms | http |
| 186.148.180.46:999 | ✓ 1056ms | ✓ 1981ms | ✓ 1532ms | ✓ 1823ms | ✓ 1522ms | http |
| 103.39.51.190:8080 | ✓ 1769ms | 否 | ✓ 1714ms | ✓ 1541ms | ✓ 1310ms | http |
| 62.113.119.14:8080 | ✓ 1166ms | ✓ 1690ms | ✓ 1491ms | ✓ 1742ms | ✓ 1263ms | http |
| 150.241.77.125:3128 | ✓ 1076ms | 否 | ✓ 1658ms | 否 | ✓ 1350ms | http |
| 103.215.36.88:13117 | 否 | ✓ 1708ms | 否 | ✓ 1688ms | ✓ 1049ms | http |
| 37.27.100.107:443 | ✓ 1872ms | 否 | 否 | ✓ 1757ms | ✓ 1954ms | http |
| 90.84.188.97:8000 | ✓ 1117ms | 否 | 否 | ✓ 1714ms | ✓ 1557ms | http |
| 103.74.192.243:7890 | 否 | 否 | ✓ 621ms | ✓ 852ms | ✓ 666ms | http |
| 46.249.103.192:443 | ✓ 980ms | 否 | ✓ 1270ms | ✓ 1620ms | 否 | http |
| 192.166.82.55:1080 | ✓ 1251ms | 否 | ✓ 1184ms | ✓ 1736ms | ✓ 752ms | http |
| 103.170.22.145:8080 | 否 | 否 | ✓ 1655ms | ✓ 1630ms | ✓ 1419ms | http |
| 172.212.68.37:3128 | ✓ 1257ms | 否 | ✓ 903ms | ✓ 1679ms | ✓ 922ms | http |
| 74.50.96.247:8888 | ✓ 489ms | ✓ 1972ms | ✓ 894ms | ✓ 979ms | 否 | http |
| 61.72.221.94:3128 | ✓ 887ms | 否 | ✓ 1933ms | ✓ 1013ms | ✓ 1772ms | http |

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
