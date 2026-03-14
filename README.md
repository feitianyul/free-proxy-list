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

最后更新：2026-03-14 07:43:55 UTC（2026-03-14 15:43:55 UTC+8）

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
| 216.180.127.45:1080 | ✓ 1104ms | 否 | ✓ 1337ms | ✓ 1586ms | ✓ 1069ms | http |
| 205.209.118.30:3138 | ✓ 1107ms | ✓ 1138ms | ✓ 832ms | ✓ 1282ms | ✓ 995ms | http |
| 85.198.96.242:3128 | ✓ 735ms | 否 | 否 | ✓ 1819ms | ✓ 1391ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1371ms | ✓ 1263ms | ✓ 991ms | http |
| 217.76.245.80:999 | ✓ 670ms | 否 | ✓ 1411ms | ✓ 1471ms | ✓ 1155ms | http |
| 45.88.0.117:3128 | ✓ 1207ms | 否 | ✓ 1897ms | ✓ 1821ms | ✓ 1728ms | http |
| 45.88.0.116:3128 | ✓ 1209ms | ✓ 1776ms | 否 | ✓ 1957ms | ✓ 1707ms | http |
| 213.220.62.62:3128 | ✓ 1208ms | 否 | ✓ 1893ms | ✓ 1821ms | ✓ 1723ms | http |
| 45.88.0.115:3128 | ✓ 1208ms | ✓ 1777ms | 否 | ✓ 1948ms | ✓ 1708ms | http |
| 45.88.0.98:3128 | ✓ 1205ms | ✓ 1738ms | 否 | ✓ 1991ms | ✓ 1721ms | http |
| 45.88.0.113:3128 | ✓ 1207ms | 否 | ✓ 1898ms | ✓ 1830ms | ✓ 1715ms | http |
| 45.88.0.99:3128 | ✓ 1207ms | ✓ 1635ms | 否 | 否 | ✓ 1805ms | http |
| 45.88.0.111:3128 | ✓ 1207ms | ✓ 1712ms | 否 | 否 | ✓ 1730ms | http |
| 45.88.0.114:3128 | ✓ 1209ms | 否 | ✓ 1895ms | ✓ 1838ms | ✓ 1705ms | http |
| 120.92.212.16:8890 | ✓ 1250ms | ✓ 1460ms | ✓ 1346ms | 否 | ✓ 1284ms | http |
| 120.92.212.16:7890 | ✓ 993ms | ✓ 1241ms | ✓ 1907ms | 否 | 否 | http |
| 150.230.249.50:1080 | ✓ 1527ms | 否 | ✓ 877ms | ✓ 1022ms | ✓ 1293ms | http |
| 81.70.169.194:80 | ✓ 1324ms | ✓ 1355ms | ✓ 1189ms | 否 | 否 | http |
| 101.43.255.96:80 | ✓ 1987ms | ✓ 1339ms | 否 | ✓ 1525ms | ✓ 1081ms | http |
| 86.53.183.16:1080 | ✓ 897ms | 否 | ✓ 1330ms | 否 | ✓ 1534ms | http |
| 43.167.227.161:1080 | 否 | 否 | ✓ 1662ms | ✓ 1039ms | ✓ 1368ms | http |
| 116.80.96.107:3172 | ✓ 1515ms | 否 | ✓ 1662ms | 否 | ✓ 1702ms | http |
| 38.145.203.135:8443 | ✓ 828ms | ✓ 1953ms | ✓ 420ms | ✓ 761ms | ✓ 590ms | http |
| 152.42.213.210:8080 | ✓ 1554ms | 否 | ✓ 1707ms | ✓ 1316ms | ✓ 1023ms | http |
| 152.42.213.210:443 | ✓ 1548ms | 否 | ✓ 1546ms | ✓ 1225ms | ✓ 1289ms | http |
| 162.240.154.26:3128 | ✓ 1106ms | 否 | ✓ 643ms | ✓ 1424ms | 否 | http |
| 210.77.29.245:7890 | ✓ 916ms | ✓ 1172ms | ✓ 1390ms | ✓ 1159ms | ✓ 918ms | http |
| 38.180.2.107:3128 | ✓ 1216ms | ✓ 1896ms | ✓ 1843ms | 否 | 否 | http |
| 45.140.147.82:1081 | ✓ 858ms | 否 | ✓ 867ms | ✓ 1647ms | ✓ 1345ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1433ms | ✓ 1796ms | ✓ 1925ms | http |
| 45.136.131.39:8443 | ✓ 909ms | ✓ 1440ms | ✓ 1473ms | ✓ 795ms | ✓ 679ms | http |
| 45.136.131.42:8447 | ✓ 907ms | ✓ 727ms | 否 | ✓ 985ms | ✓ 673ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1927ms | 否 | ✓ 1492ms | ✓ 1163ms | http |
| 114.214.208.153:10808 | ✓ 1458ms | 否 | ✓ 1644ms | ✓ 1699ms | ✓ 1457ms | http |
| 45.140.147.82:1082 | ✓ 565ms | ✓ 1795ms | ✓ 1657ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1250ms | ✓ 1363ms | ✓ 1621ms | 否 | ✓ 1106ms | http |
| 35.225.22.61:80 | ✓ 1204ms | ✓ 1379ms | 否 | ✓ 1031ms | 否 | http |
| 167.71.196.28:8080 | 否 | 否 | ✓ 751ms | ✓ 1086ms | ✓ 845ms | http |
| 45.136.198.40:3128 | ✓ 1218ms | 否 | ✓ 1058ms | ✓ 1609ms | ✓ 1312ms | http |
| 101.32.244.83:8080 | ✓ 957ms | 否 | ✓ 974ms | ✓ 1497ms | ✓ 1222ms | http |
| 121.43.196.210:8222 | ✓ 915ms | ✓ 1062ms | ✓ 881ms | ✓ 1180ms | ✓ 924ms | http |
| 121.43.196.213:8222 | ✓ 902ms | ✓ 1129ms | ✓ 864ms | ✓ 1151ms | ✓ 951ms | http |
| 114.55.226.123:10086 | ✓ 1055ms | ✓ 1490ms | ✓ 1025ms | ✓ 1289ms | ✓ 1075ms | http |
| 24.144.86.173:1080 | ✓ 1077ms | ✓ 1431ms | ✓ 956ms | ✓ 825ms | ✓ 742ms | http |
| 1.225.116.115:1080 | 否 | 否 | ✓ 1400ms | ✓ 1179ms | ✓ 1036ms | http |
| 150.249.255.91:3128 | ✓ 1617ms | ✓ 1045ms | ✓ 542ms | ✓ 1918ms | ✓ 1746ms | http |
| 34.101.184.164:3128 | ✓ 1760ms | 否 | 否 | ✓ 1470ms | ✓ 1227ms | http |
| 106.117.208.101:7890 | ✓ 1513ms | 否 | 否 | ✓ 1455ms | ✓ 1094ms | http |
| 101.47.73.135:3128 | ✓ 1931ms | 否 | ✓ 1960ms | ✓ 1381ms | 否 | http |
| 202.155.12.161:443 | ✓ 1345ms | ✓ 1896ms | ✓ 924ms | ✓ 1189ms | ✓ 933ms | http |
| 123.57.0.163:8888 | ✓ 1558ms | ✓ 1509ms | 否 | 否 | ✓ 1498ms | http |
| 172.212.68.37:3128 | ✓ 683ms | ✓ 1655ms | ✓ 1156ms | ✓ 1438ms | ✓ 1241ms | http |
| 16.78.119.130:443 | 否 | ✓ 1771ms | ✓ 1651ms | 否 | ✓ 1622ms | http |
| 43.165.195.107:3128 | ✓ 1530ms | 否 | ✓ 974ms | ✓ 1223ms | 否 | http |
| 120.92.211.211:7890 | ✓ 996ms | 否 | ✓ 1882ms | 否 | ✓ 972ms | http |
| 121.40.231.103:7890 | ✓ 1087ms | ✓ 1065ms | ✓ 928ms | ✓ 1218ms | ✓ 870ms | http |
| 103.39.51.190:8080 | ✓ 1917ms | 否 | 否 | ✓ 1502ms | ✓ 1400ms | http |
| 61.52.131.172:8443 | ✓ 908ms | 否 | 否 | ✓ 1172ms | ✓ 935ms | http |
| 62.113.119.14:8080 | ✓ 744ms | 否 | ✓ 1074ms | ✓ 1759ms | ✓ 1206ms | http |
| 193.168.173.136:443 | ✓ 967ms | 否 | 否 | ✓ 1941ms | ✓ 1780ms | http |
| 139.159.99.242:8080 | ✓ 838ms | ✓ 1015ms | ✓ 994ms | 否 | ✓ 868ms | http |
| 162.248.165.72:1080 | 否 | 否 | ✓ 1502ms | ✓ 1776ms | ✓ 1822ms | http |

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
