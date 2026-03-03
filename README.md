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

最后更新：2026-03-03 15:44:16 UTC（2026-03-03 23:44:16 UTC+8）

**代理总数：41**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 41 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 41 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 162.240.154.26:3128 | ✓ 704ms | ✓ 1568ms | 否 | ✓ 1595ms | ✓ 1315ms | http |
| 166.0.192.117:8888 | 否 | ✓ 1444ms | ✓ 973ms | 否 | ✓ 1064ms | http |
| 205.209.118.30:3138 | ✓ 852ms | 否 | 否 | ✓ 1363ms | ✓ 1328ms | http |
| 217.76.245.80:999 | ✓ 868ms | ✓ 1625ms | ✓ 1893ms | ✓ 1581ms | ✓ 1309ms | http |
| 186.148.180.46:999 | ✓ 1231ms | 否 | ✓ 1435ms | ✓ 1775ms | ✓ 1463ms | http |
| 35.234.17.221:8080 | ✓ 887ms | ✓ 1750ms | ✓ 1215ms | 否 | 否 | http |
| 81.70.169.194:80 | ✓ 1046ms | 否 | ✓ 1055ms | ✓ 1607ms | 否 | http |
| 101.43.255.96:80 | 否 | 否 | ✓ 1287ms | ✓ 1391ms | ✓ 1080ms | http |
| 62.113.119.14:8080 | ✓ 784ms | 否 | ✓ 1331ms | 否 | ✓ 1847ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1257ms | 否 | ✓ 1219ms | ✓ 976ms | http |
| 91.193.240.157:9877 | ✓ 924ms | 否 | ✓ 1003ms | ✓ 1989ms | ✓ 1764ms | http |
| 115.231.181.40:8128 | ✓ 1397ms | 否 | ✓ 1309ms | 否 | ✓ 878ms | http |
| 115.76.5.32:10009 | 否 | 否 | ✓ 1441ms | ✓ 1968ms | ✓ 1474ms | http |
| 192.71.213.85:9090 | ✓ 1594ms | 否 | ✓ 1786ms | ✓ 1922ms | 否 | http |
| 103.84.95.54:7890 | ✓ 1305ms | 否 | ✓ 1595ms | 否 | ✓ 944ms | http |
| 45.140.147.155:1081 | 否 | 否 | ✓ 1146ms | ✓ 1742ms | ✓ 1251ms | http |
| 142.171.85.32:1080 | ✓ 600ms | ✓ 1139ms | ✓ 1261ms | ✓ 844ms | ✓ 1442ms | http |
| 5.75.196.26:40000 | ✓ 1035ms | ✓ 1956ms | 否 | ✓ 1824ms | ✓ 1257ms | http |
| 120.92.212.16:8890 | ✓ 1776ms | 否 | ✓ 1033ms | 否 | ✓ 1006ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1346ms | ✓ 1922ms | 否 | ✓ 1329ms | http |
| 221.127.195.224:8888 | ✓ 1313ms | 否 | ✓ 1179ms | ✓ 1320ms | ✓ 1225ms | http |
| 132.226.235.199:1080 | ✓ 1729ms | 否 | 否 | ✓ 1290ms | ✓ 989ms | http |
| 119.134.178.238:7890 | 否 | 否 | ✓ 1596ms | ✓ 1626ms | ✓ 1370ms | http |
| 101.32.244.83:8080 | ✓ 1396ms | 否 | ✓ 1009ms | ✓ 1363ms | ✓ 1059ms | http |
| 121.43.196.213:8222 | ✓ 941ms | ✓ 1109ms | ✓ 884ms | ✓ 1138ms | ✓ 919ms | http |
| 121.43.196.210:8222 | ✓ 873ms | ✓ 1123ms | ✓ 949ms | ✓ 1139ms | ✓ 948ms | http |
| 114.55.226.123:10086 | ✓ 1083ms | ✓ 1696ms | ✓ 1007ms | ✓ 1282ms | ✓ 1049ms | http |
| 188.166.208.168:9876 | 否 | 否 | ✓ 899ms | ✓ 1173ms | ✓ 855ms | http |
| 45.136.198.40:3128 | ✓ 758ms | 否 | ✓ 729ms | ✓ 1594ms | ✓ 1321ms | http |
| 2.56.178.131:443 | ✓ 950ms | 否 | ✓ 1888ms | ✓ 1918ms | 否 | http |
| 103.82.23.118:5247 | ✓ 1761ms | 否 | ✓ 1349ms | 否 | ✓ 1275ms | http |
| 103.39.51.190:8080 | ✓ 1824ms | 否 | 否 | ✓ 1383ms | ✓ 1395ms | http |
| 194.59.204.87:9080 | ✓ 1063ms | ✓ 1770ms | ✓ 1093ms | 否 | 否 | http |
| 45.88.0.99:3128 | ✓ 1507ms | 否 | ✓ 1759ms | 否 | ✓ 1791ms | http |
| 45.88.0.115:3128 | ✓ 680ms | ✓ 1774ms | 否 | ✓ 1790ms | 否 | http |
| 160.238.65.8:3128 | ✓ 1683ms | 否 | ✓ 874ms | ✓ 1490ms | 否 | http |
| 45.88.0.98:3128 | ✓ 1095ms | 否 | ✓ 1124ms | 否 | ✓ 1863ms | http |
| 77.83.203.6:443 | 否 | ✓ 1851ms | ✓ 1529ms | 否 | ✓ 1899ms | http |
| 45.10.69.102:8888 | ✓ 425ms | 否 | ✓ 313ms | ✓ 777ms | ✓ 1027ms | http |
| 121.230.8.220:1080 | ✓ 1568ms | ✓ 1566ms | ✓ 1212ms | ✓ 1442ms | 否 | http |
| 47.95.231.180:8084 | ✓ 1122ms | ✓ 1540ms | ✓ 1568ms | ✓ 1795ms | ✓ 973ms | http |

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
