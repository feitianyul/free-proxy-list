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

最后更新：2026-03-10 16:59:47 UTC（2026-03-11 00:59:47 UTC+8）

**代理总数：49**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 154.3.236.202:3128 | ✓ 1122ms | 否 | ✓ 1329ms | ✓ 1481ms | ✓ 1137ms | http |
| 202.155.12.161:443 | ✓ 1847ms | 否 | 否 | ✓ 1940ms | ✓ 1562ms | http |
| 205.209.118.30:3138 | ✓ 1690ms | 否 | ✓ 1570ms | ✓ 1795ms | 否 | http |
| 45.140.147.155:1081 | ✓ 1069ms | 否 | ✓ 617ms | ✓ 1561ms | ✓ 1232ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 718ms | ✓ 917ms | ✓ 803ms | http |
| 185.191.236.162:3128 | ✓ 1109ms | 否 | ✓ 1338ms | ✓ 1837ms | ✓ 1298ms | http |
| 45.136.131.47:8443 | 否 | 否 | ✓ 769ms | ✓ 1126ms | ✓ 1043ms | http |
| 190.212.131.238:3128 | ✓ 1845ms | 否 | ✓ 1765ms | 否 | ✓ 1579ms | http |
| 190.9.109.198:999 | ✓ 1068ms | 否 | ✓ 1328ms | ✓ 1433ms | ✓ 1285ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1553ms | ✓ 1186ms | 否 | ✓ 1071ms | http |
| 67.169.98.211:443 | ✓ 1604ms | 否 | 否 | ✓ 1905ms | ✓ 704ms | http |
| 81.70.169.194:80 | ✓ 1239ms | 否 | 否 | ✓ 1178ms | ✓ 1977ms | http |
| 35.225.22.61:80 | ✓ 448ms | 否 | ✓ 434ms | ✓ 1200ms | ✓ 928ms | http |
| 45.136.198.40:3128 | ✓ 1987ms | ✓ 1772ms | ✓ 1950ms | 否 | ✓ 1770ms | http |
| 192.227.137.63:5050 | ✓ 1303ms | 否 | ✓ 331ms | ✓ 1796ms | 否 | http |
| 120.92.212.16:8890 | ✓ 904ms | ✓ 1203ms | 否 | ✓ 1193ms | 否 | http |
| 39.104.201.40:7890 | ✓ 945ms | ✓ 1135ms | ✓ 997ms | ✓ 1160ms | ✓ 919ms | http |
| 158.69.185.37:3129 | ✓ 664ms | 否 | 否 | ✓ 1505ms | ✓ 1410ms | http |
| 162.240.154.26:3128 | ✓ 948ms | 否 | 否 | ✓ 1877ms | ✓ 1564ms | http |
| 150.107.140.238:3128 | ✓ 1902ms | 否 | ✓ 968ms | 否 | ✓ 888ms | http |
| 152.42.213.210:8080 | ✓ 726ms | 否 | 否 | ✓ 1157ms | ✓ 1433ms | http |
| 101.43.255.96:80 | 否 | ✓ 1766ms | ✓ 1200ms | 否 | ✓ 1276ms | http |
| 106.14.203.63:3333 | ✓ 803ms | 否 | ✓ 896ms | ✓ 1040ms | ✓ 838ms | http |
| 120.92.212.16:7890 | ✓ 924ms | ✓ 1170ms | 否 | 否 | ✓ 907ms | http |
| 45.136.130.223:8443 | ✓ 419ms | ✓ 739ms | ✓ 749ms | ✓ 764ms | ✓ 571ms | http |
| 46.183.25.8:443 | ✓ 891ms | 否 | 否 | ✓ 1556ms | ✓ 799ms | http |
| 86.53.183.16:1080 | ✓ 643ms | ✓ 1936ms | ✓ 1622ms | ✓ 1939ms | 否 | http |
| 192.227.137.65:5050 | 否 | 否 | ✓ 1023ms | ✓ 1586ms | ✓ 1885ms | http |
| 138.124.53.25:7443 | ✓ 628ms | 否 | ✓ 1613ms | ✓ 1577ms | 否 | http |
| 115.231.181.40:8128 | ✓ 857ms | 否 | ✓ 1104ms | 否 | ✓ 1126ms | http |
| 168.235.110.63:3128 | 否 | 否 | ✓ 1513ms | ✓ 1347ms | ✓ 982ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1838ms | ✓ 1638ms | ✓ 1460ms | http |
| 47.77.193.180:1080 | ✓ 131ms | ✓ 1799ms | ✓ 211ms | ✓ 753ms | ✓ 532ms | http |
| 101.47.73.135:3128 | ✓ 1428ms | 否 | 否 | ✓ 1837ms | ✓ 1701ms | http |
| 1.231.81.166:3128 | ✓ 1460ms | ✓ 1038ms | ✓ 1218ms | ✓ 848ms | ✓ 697ms | http |
| 165.227.5.10:8888 | ✓ 604ms | ✓ 878ms | 否 | 否 | ✓ 542ms | http |
| 152.70.98.46:8888 | ✓ 1124ms | 否 | ✓ 990ms | ✓ 853ms | ✓ 809ms | http |
| 217.77.102.18:3128 | ✓ 1439ms | 否 | ✓ 1577ms | 否 | ✓ 1720ms | http |
| 89.185.85.138:1080 | ✓ 1906ms | ✓ 1896ms | ✓ 1816ms | ✓ 1724ms | ✓ 1423ms | http |
| 150.249.255.91:3128 | ✓ 1702ms | 否 | ✓ 1174ms | 否 | ✓ 892ms | http |
| 116.80.96.102:3172 | ✓ 1471ms | 否 | ✓ 1447ms | ✓ 1877ms | 否 | http |
| 43.167.227.161:1080 | 否 | ✓ 1869ms | ✓ 837ms | 否 | ✓ 1932ms | http |
| 103.39.51.190:8080 | ✓ 1661ms | 否 | 否 | ✓ 1454ms | ✓ 1348ms | http |
| 194.213.18.200:443 | ✓ 753ms | 否 | ✓ 928ms | ✓ 1305ms | 否 | http |
| 91.107.141.42:8081 | 否 | 否 | ✓ 1311ms | ✓ 1983ms | ✓ 1965ms | http |
| 61.52.131.172:8443 | ✓ 866ms | ✓ 1118ms | ✓ 980ms | ✓ 1208ms | ✓ 901ms | http |
| 221.122.91.36:11195 | 否 | ✓ 1207ms | ✓ 873ms | ✓ 1201ms | ✓ 946ms | http |
| 221.122.91.36:11273 | 否 | ✓ 1185ms | ✓ 969ms | ✓ 1119ms | ✓ 936ms | http |
| 103.113.70.189:1081 | ✓ 382ms | ✓ 1344ms | 否 | ✓ 1523ms | ✓ 979ms | http |

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
